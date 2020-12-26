package util

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenDB() (db *sqlx.DB, errs error) {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/userdb")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		errs = err
		return
	}
	db = database

	return
}
