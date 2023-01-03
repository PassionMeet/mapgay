package db

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@127.0.0.1/jipeng")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

const (
	// tables
	tableUserGeo = "user_geo"
)

func InsertUserGeo(ctx context.Context) {
	sq.Insert(tableUserGeo).Columns("")
}
