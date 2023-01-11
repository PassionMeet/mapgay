package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlCli *sql.DB

func InitMySQL() {
	var err error
	mysqlCli, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/jipeng")
	if err != nil {
		panic(err)
	}
	mysqlCli.SetConnMaxLifetime(time.Minute * 3)
	mysqlCli.SetMaxOpenConns(10)
	mysqlCli.SetMaxIdleConns(10)
	err = mysqlCli.Ping()
	if err != nil {
		panic(err)
	}
}
