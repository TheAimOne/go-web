package database_util

import (
	"fmt"

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
	whereCondition := " "

	for _, criteria := range filter.Criterias {
		key, keyExists := filterMap[criteria.Key]
		if keyExists {
			switch criteria.Operator {
			case model.EQUALS:
				whereCondition += fmt.Sprintf(" %s = %s ", key, criteria.Value)
			case model.IN:
				whereCondition += fmt.Sprintf(" %s in (%s) ", key, criteria.Values)
			case model.LIKE:
				whereCondition += fmt.Sprintf(" %s LIKE "+"'%%"+"%s"+"%%'", key, criteria.Value)
			}
		}

	}
	return whereCondition
}
