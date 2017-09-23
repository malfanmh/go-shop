package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB(endpoint string) (err error) {
	db, err = sql.Open("mysql", endpoint)
	return
}
