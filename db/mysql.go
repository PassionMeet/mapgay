package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/cmfunc/jipeng/conf"
	_ "github.com/go-sql-driver/mysql"
)

var mysqlCli *sql.DB

func InitMySQL(cfg *conf.MySQL) {
	var err error
	var uri = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	mysqlCli, err = sql.Open("mysql", uri)
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
