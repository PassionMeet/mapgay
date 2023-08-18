package db

import (
	"context"

	"github.com/pkg/errors"
)

type InsertAStrongholdParam struct {
	Name            string
	Intro           string
	PublicNotice    string
	LocationPolygon string
	LocationPoint   string
	LocationLine    *string //use nil to transfor null
	DetailAddress   string
	HotLevel        string
}

func InsertAStronghold(ctx context.Context, param *InsertAStrongholdParam) error {
	query := `insert into stronghold 
	(name,intro,public_notice,location_polygon,location_point,location_line,detail_address,hot_level) 
	values 
	(?,?,?,?,?,?,?,?)`
	_, err := mysqlCli.ExecContext(ctx, query)
	if err != nil {
		return errors.Wrapf(err, "InsertAStronghold param:%+v", param)
	}
	return nil
}
