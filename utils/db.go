package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	// fmt.Println("hello")
	Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/keshe")
	if err != nil {
		panic(err.Error())
	}
}
