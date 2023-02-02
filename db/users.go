package db

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
)

const (
	UsersTable = "users" // 用户表
)

type UsersRow struct {
	ID           uint64
	Openid       string
	SessionKey   string
	RegisterTime time.Time
	Username     string
	Avatar       string
	Height       uint32
	Weight       uint32
	Age          uint32
	Length       uint32
	WeixinID     string
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

func GetUser(ctx context.Context, openid string) (UsersRow, error) {
	row := UsersRow{}
	err := sq.Select("username,avatar,height,weight,age,length,weixin_id").
		From(UsersTable).Where(sq.Eq{"openid": openid}).
		RunWith(mysqlCli).QueryRowContext(ctx).
		Scan(&row.Username, &row.Avatar, &row.Height, &row.Weight, &row.Age, &row.Length, &row.WeixinID)
	return row, err
}
