package db

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

const (
	UsersTable = "users" // 用户表
)

type UsersRow struct {
	Openid     string
	SessionKey string
}

func InsertUsers(ctx context.Context, rows *UsersRow) (sql.Result, error) {
	return sq.Insert(UsersTable).
		Columns("openid,session_key").
		Values(rows.Openid, rows.SessionKey).
		RunWith(db).
		ExecContext(ctx)
}
