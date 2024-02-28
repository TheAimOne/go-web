package function

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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
	SelectPaginateAndFilterByQuery(query string, filter model.Filter, filterMap map[string]string) (*sql.Rows, error)
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

func GetQueryByFilter(table string, filter model.Filter, columns []string,
	filterMap map[string]string) (string, error) {
	columnString, _, err := database_util.ColumnHelper(columns)
	if err != nil {
		return "", constants.ErrorCreatingSql
	}

	finalQuery := fmt.Sprintf("select %s from %s %s %s %s",
		columnString,
		table,
		database_util.AddWhereCondition(filterMap, &filter, true),
		database_util.SortingHelper(filter),
		database_util.PaginationHelper(filter))
	return finalQuery, nil
}

func (f *functionImpl) SelectPaginateAndFilter(table string, filter model.Filter, columns []string,
	filterMap map[string]string) (*sql.Rows, error) {

	finalQuery, err := GetQueryByFilter(table, filter, columns, filterMap)
	if err != nil {
		return nil, err
	}

	rows, err := connection.DB.Query(finalQuery)

	return rows, err
}

func (f *functionImpl) SelectPaginateAndFilterByQuery(query string, filter model.Filter, filterMap map[string]string) (*sql.Rows, error) {
	addWhereCondition := true
	if strings.Contains(strings.ToLower(query), "where") {
		addWhereCondition = false
		query = query + " AND "
	}

	finalQuery := fmt.Sprintf(" %s %s %s %s",
		query,
		" ( "+database_util.AddWhereCondition(filterMap, &filter, addWhereCondition)+" ) ",
		database_util.SortingHelper(filter),
		database_util.PaginationHelper(filter))
	log.Println("finalQuery", finalQuery)
	rows, err := connection.DB.Query(finalQuery)
	return rows, err
}
