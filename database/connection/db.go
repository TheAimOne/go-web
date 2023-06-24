package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const FUR = "FUR"

const (
	host     = "localhost"
	port     = 5432
	user     = "api"
	password = "password"
	dbname   = "database_dummy"
)

var DB *sql.DB
var S *string

type Test struct {
	DB1 *sql.DB
}

var T Test

func InitDB() {
	fmt.Println("dinga")
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	var err error
	DB, err = sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = DB.Ping()
	CheckError(err)

	fmt.Println("Connected!", DB)

	s := "fasfa"

	T := &Test{DB}

	fmt.Println("Connected!", T)

	S = &s

}

func GetDB() *sql.DB {
	return DB
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
