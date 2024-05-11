package database_util

import (
	"fmt"
	"strconv"
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

func AddWhereCondition(filterMap map[string]string, filter *model.Filter, addWhereKeyword bool) string {
	whereCondition := make([]string, 0)
	whereKeyword := ""
	if addWhereKeyword {
		whereKeyword = " WHERE "
	}
	if len(filter.Criterias) == 0 {
		return whereKeyword + " 1=1 "
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
					fmt.Sprintf(" UPPER(%s) LIKE UPPER("+"'%%"+"%s"+"%%')", columnName, criteria.Value))
			case model.NOT_EQUALS:
				whereCondition = append(whereCondition,
					fmt.Sprintf(" %s <> '%s' ", columnName, criteria.Value))
			case model.GREATER_THAN:
				whereCondition = append(whereCondition, fmt.Sprintf(" %s > '%s' ", columnName, criteria.Value))
			case model.LESS_THAN:
				whereCondition = append(whereCondition, fmt.Sprintf(" %s < '%s' ", columnName, criteria.Value))
			}
		}

	}

	if filter.IsAnd {
		return whereKeyword + strings.Join(whereCondition, " AND ")
	}
	return whereKeyword + strings.Join(whereCondition, " OR ")
}

func PaginationHelper(filter model.Filter) string {
	pageString := ""
	if filter.PageSize != 0 {
		offset := filter.PageNumber * filter.PageSize
		pageString = fmt.Sprintf("limit %s offset %s", strconv.FormatInt(filter.PageSize, 10),
			strconv.FormatInt(offset, 10))

	}
	return pageString
}

func SortingHelper(filter model.Filter) string {
	sortString := ""
	if filter.SortKey != "" {
		if filter.SortDirection == "" {
			filter.SortDirection = "asc"
		}
		sortString = fmt.Sprintf("order by %s %s", filter.SortKey, filter.SortDirection)
	}
	return sortString
}
