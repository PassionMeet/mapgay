package db

import (
	"context"

	"github.com/pkg/errors"
)

type UmsAccount struct {
	ID       int64
	Account  string
	Email    string
	Password string
	Role     string
}

func GetUmsAccount(ctx context.Context, account string) (*UmsAccount, error) {
	resp := UmsAccount{}
	query := `select id,account,email,password,role from ums_account where account = ?`
	err := mysqlCli.QueryRowContext(ctx, query, account).Scan(
		&resp.ID, &resp.Account, &resp.Email, &resp.Password, &resp.Role,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "GetUmsAccount account %s", account)
	}
	return &resp, nil
}
