package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
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
	query := "INSERT INTO users (openid,session_key) VALUES (?,?) ON DUPLICATE KEY UPDATE session_key=values(session_key)"
	return mysqlCli.ExecContext(ctx, query, rows.Openid, rows.SessionKey)
}

func UpdateUser(ctx context.Context, where interface{}, update sq.Eq) error {
	_, err := sq.Update(UsersTable).SetMap(update).Where(where).RunWith(mysqlCli).ExecContext(ctx)
	if err != nil {
		err = errors.Wrapf(err, "UpdateUser where:%+v update:%+v", where, update)
		return err
	}
	return nil
}

func GetUser(ctx context.Context, openid string) (*UsersRow, error) {
	row := &UsersRow{}
	query := `select username,avatar,height,weight,age,length,weixin_id from users where openid = ?`
	err := mysqlCli.QueryRowContext(ctx, query, openid).Scan(&row.Username, &row.Avatar, &row.Height, &row.Weight, &row.Age, &row.Length, &row.WeixinID)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func GetUsers(ctx context.Context, openids []string) (users map[string]*UsersRow, err error) {
	for i, openid := range openids {
		openids[i] = fmt.Sprintf("'%s'", openid)
	}
	query := fmt.Sprintf(`select openid,username,avatar,height,weight,age,length,weixin_id from users where openid in (%s)`, strings.Join(openids, ","))
	log.Printf("GetUsers opendis:%v query:%s", openids, query)
	rows, err := mysqlCli.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	users = make(map[string]*UsersRow, 0)
	for rows.Next() {
		row := &UsersRow{}
		err = rows.Scan(&row.Openid, &row.Username, &row.Avatar, &row.Height, &row.Weight, &row.Age, &row.Length, &row.WeixinID)
		if err != nil {
			return nil, err
		}
		users[row.Openid] = row
	}
	return users, nil
}
