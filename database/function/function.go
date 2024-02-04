package function

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/go-web/database/connection"
	database_util "github.com/go-web/database/util"
	"github.com/go-web/pkg/constants"
	"github.com/go-web/pkg/model"
)

type DBFunction interface {
	Insert(table string, columns []string, values []interface{}) error
	SelectAll(table, condition string, columns []string) (*sql.Rows, error)
	Select(table, condition string, columns []string) (*sql.Row, error)
	SelectRaw(query string) (*sql.Rows, error)
	SelectPaginateAndFilter(table string, filter model.Filter, columns []string, filterMap map[string]string) (*sql.Rows, error)
}

func NewDBFunction() DBFunction {
	return &functionImpl{}
}

type functionImpl struct {
}

func (f *functionImpl) Insert(table string, columns []string, values []interface{}) error {
	columnString, valueString, err := database_util.ColumnHelper(columns)

	if err != nil || len(columns) != len(values) {
		return constants.ErrorCreatingSql
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columnString, valueString)

	log.Println(sql)
	result, err := connection.DB.Exec(sql, values...)

	log.Println(result)

	return err
}

func (f *functionImpl) SelectAll(table, condition string, columns []string) (*sql.Rows, error) {
	columnString, _, err := database_util.ColumnHelper(columns)
	if err != nil {
		return nil, constants.ErrorCreatingSql
	}
	query := fmt.Sprintf("select %s from %s %s", columnString, table, condition)

	rows, err := connection.DB.Query(query)

	return rows, nil
}

func (f *functionImpl) Select(table, condition string, columns []string) (*sql.Row, error) {
	columnString, _, err := database_util.ColumnHelper(columns)
	if err != nil {
		log.Println(err)
		return nil, constants.ErrorCreatingSql
	}
	log.Println(fmt.Sprintf("select %s from %s %s", columnString, table, condition))

	rows := connection.DB.QueryRow(fmt.Sprintf("select %s from %s %s", columnString, table, condition))

	if err != nil {
		log.Println(err)
		return nil, constants.ErrorReadingFromDB
	}
	if rows == nil {
		return nil, constants.ErrorNoRecordsInDB
	}

	return rows, nil
}

func (f *functionImpl) SelectRaw(query string) (*sql.Rows, error) {
	rows, err := connection.DB.Query(query)

	return rows, err
}

func (f *functionImpl) SelectPaginateAndFilter(table string, filter model.Filter, columns []string,
	filterMap map[string]string) (*sql.Rows, error) {
	columnString, _, err := database_util.ColumnHelper(columns)
	if err != nil {
		return nil, constants.ErrorCreatingSql
	}

	pageString := ""
	if filter.PageSize != 0 {
		offset := filter.PageNumber * filter.PageSize
		pageString = fmt.Sprintf("limit %s offset %s", strconv.FormatInt(filter.PageSize, 10),
			strconv.FormatInt(offset, 10))

	}

	sortString := ""

	if filter.SortKey != "" {
		if filter.SortDirection == "" {
			filter.SortDirection = "asc"
		}
		sortString = fmt.Sprintf("order by %s %s", filter.SortKey, filter.SortDirection)
	}

	whereCondition := database_util.AddWhereCondition(filterMap, &filter)
	finalQuery := fmt.Sprintf("select %s from %s %s %s %s", columnString, table, whereCondition, sortString, pageString)
	log.Println(finalQuery)
	rows, err := connection.DB.Query(finalQuery)

	return rows, err
}
