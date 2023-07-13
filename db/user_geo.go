package db

import "context"

// 用户的位置信息
type UserGeo struct {
	Openid    string
	Latitude  float64 `json:"latitude"`  //纬度
	Longitude float64 `json:"longitude"` //经度
}

func InsertUserGeo(ctx context.Context, row *UserGeo) error {
	return nil
}
