package function

import (
	"fmt"
	"log"

	"github.com/go-web/database/connection"
	"github.com/go-web/pkg/constants"
)

type DBFunction interface {
	Insert(table string, columns []string, values []interface{}) error
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
