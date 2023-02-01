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
	query, args, err := sq.Insert(UsersTable).
		Columns("openid,session_key").
		Values(rows.Openid, rows.SessionKey).
		ToSql()
	if err != nil {
		return nil, err
	}
	query = query + " ON DUPLICATE KEY UPDATE openid=values(openid), session_key=values(session_key)"
	return mysqlCli.ExecContext(ctx, query, args...)
}

func UpdateUser(ctx context.Context, where interface{}, update sq.Eq) error {
	_, err := sq.Update(UsersTable).SetMap(update).Where(where).RunWith(mysqlCli).ExecContext(ctx)
	return err
}
