package function

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-web/database/connection"
	"github.com/go-web/pkg/constants"
)

type DBFunction interface {
	Insert(table string, columns []string, values []interface{}) error
	SelectAll(table, condition string, columns []string) (*sql.Rows, error)
	Select(table, condition string, columns []string) (*sql.Row, error)
	SelectRaw(query string) (*sql.Rows, error)
}

func NewDBFunction() DBFunction {
	return &functionImpl{}
}

type functionImpl struct {
}

func (f *functionImpl) Insert(table string, columns []string, values []interface{}) error {
	columnString, valueString, err := columnHelper(columns)

	if err != nil || len(columns) != len(values) {
		return constants.ErrorCreatingSql
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columnString, valueString)

	result, err := connection.DB.Exec(sql, values...)

	log.Println(result)

	return err
}

func columnHelper(columns []string) (string, string, error) {
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

func (f *functionImpl) SelectAll(table, condition string, columns []string) (*sql.Rows, error) {
	columnString, _, err := columnHelper(columns)
	if err != nil {
		return nil, constants.ErrorCreatingSql
	}

	rows, err := connection.DB.Query(fmt.Sprintf("select %s from %s %s", columnString, table, condition))

	return rows, nil
}

func (f *functionImpl) Select(table, condition string, columns []string) (*sql.Row, error) {
	columnString, _, err := columnHelper(columns)
	if err != nil {
		return nil, constants.ErrorCreatingSql
	}

	rows := connection.DB.QueryRow(fmt.Sprintf("select %s from %s %s", columnString, table, condition))

	return rows, nil
}

func (f *functionImpl) SelectRaw(query string) (*sql.Rows, error) {
	rows, err := connection.DB.Query(query)

	return rows, err
}
