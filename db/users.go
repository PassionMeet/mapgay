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
	query := "INSERT INTO users (openid,session_key) VALUES (?,?) ON DUPLICATE KEY UPDATE openid=values(openid), session_key=values(session_key)"
	return mysqlCli.ExecContext(ctx, query, rows.Openid, rows.SessionKey)
}

func UpdateUser(ctx context.Context, where interface{}, update sq.Eq) error {
	_, err := sq.Update(UsersTable).SetMap(update).Where(where).RunWith(mysqlCli).ExecContext(ctx)
	return err
}

func GetUser(ctx context.Context, openid string) (*UsersRow, error) {
	row := &UsersRow{}
	query := `select username,avatar,height,weight,age,length,weixin_id from users where openid = ?`
	err := mysqlCli.QueryRowContext(ctx, query, openid).Scan(row.Username, row.Avatar, row.Height, row.Weight, row.Age, row.Length, row.WeixinID)
	if err != nil {
		return nil, err
	}
	return row, nil
}
