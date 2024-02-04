package database_util

import (
	"fmt"
	"strings"

	"github.com/go-web/pkg/model"
)

func ColumnHelper(columns []string) (string, string, error) {
	column := ""
	values := ""
	for i := 0; i < len(columns); i++ {
		if i == 0 {
			column += columns[i]
			values += fmt.Sprintf("$%d", i+1)
		} else {
			column = column + "," + columns[i]
			values = values + fmt.Sprintf(",$%d", i+1)
		}
	}

	return column, values, nil
}

func AddWhereCondition(filterMap map[string]string, filter *model.Filter) string {
	whereCondition := make([]string, 0)
	if len(filter.Criterias) == 0 {
		return ""
	}

	for _, criteria := range filter.Criterias {
		columnName, keyExists := filterMap[criteria.Key]
		if keyExists {
			switch criteria.Operator {
			case model.EQUALS:
				whereCondition = append(whereCondition,
					fmt.Sprintf(" %s = '%s' ", columnName, criteria.Value))
			case model.IN:
				if len(criteria.Values) != 0 {
					commaSeparatedValues := strings.Join(criteria.Values, ",")
					whereCondition = append(whereCondition,
						fmt.Sprintf(" %s in (%s) ", columnName, commaSeparatedValues))
				}
			case model.CONTAINS:
				whereCondition = append(whereCondition,
					fmt.Sprintf(" %s LIKE "+"'%%"+"%s"+"%%'", columnName, criteria.Value))
			case model.NOT_EQUALS:
				whereCondition = append(whereCondition,
					fmt.Sprintf(" %s <> '%s' ", columnName, criteria.Value))
			}
		}

	}

	if filter.IsAnd {
		return " WHERE " + strings.Join(whereCondition, " AND ")
	}
	return " WHERE " + strings.Join(whereCondition, " OR ")
}
