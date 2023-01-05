package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMySQL() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/jipeng")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
