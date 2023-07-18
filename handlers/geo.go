package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cmfunc/jipeng/cache"
	"github.com/cmfunc/jipeng/db"
	"github.com/gin-gonic/gin"
)

type UploadGeoRequest struct {
	Openid             string  `json:"openid"`
	Latitude           float64 `json:"latitude"`  //纬度
	Longitude          float64 `json:"longitude"` //经度
	Speed              int32   `json:"speed"`     //速度m/s
	Accuracy           float64 `json:"accuracy"`  //位置精确度
	Altitude           float64 `json:"altitude"`  //高度m
	ErrMsg             string  `json:"err_msg"`
	VerticalAccuracy   float64 `json:"verticalAccuracy"`   //垂直精度m
	HorizontalAccuracy float64 `json:"horizontalAccuracy"` //水平精度
}

func UploadGeo(ctx *gin.Context) {
	param := UploadGeoRequest{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		log.Printf("UploadGeo param:%+v %s", param, err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	param.Openid = ctx.GetString("openid")
	// 位置信息发送到nsq
	// nsq消费者单独处理
	log.Printf("UploadGeo param:%+v", param)

	// 保存位置信息
	err = cache.AddUserGeo(ctx, &cache.UserGeo{
		Openid:    param.Openid,
		Latitude:  param.Latitude,
		Longitude: param.Longitude,
	})
	if err != nil {
		log.Printf("UploadGeo param:%+v %s", param, err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, nil)

}

type GetUsersByGeoParam struct {
	Latitude  float64 `form:"latitude"`
	Longitude float64 `form:"longitude"`
	Distance  float64 `form:"distance"` //中心点，多少米范围内
}

type GetUsersByGeoResp_Item struct {
	// 和微信小程序map组建里的marker保持相似
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Avatar    string  `json:"avatar"`
	Feature   string  `json:"feature"`
	WeixinID  string  `json:"weixinID"`
}

type GetUsersByGeoResp struct {
	List []*GetUsersByGeoResp_Item `json:"list"`
}

func GetUsersByGeo(ctx *gin.Context) {
	param := GetUsersByGeoParam{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		log.Printf("GetUsersByGeo param:%+v %s", param, err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	openid := ctx.GetString("openid")
	// 通过geo信息筛选出当前地图中所有用户
	// 时间筛选
	filter := &cache.GeoFilter{
		Openid: openid,
	}
	usergeos, err := cache.GetUsersByGeo(ctx, filter)
	if err != nil {
		log.Printf("GetUsersByGeo cache.GetUsersByGeo filter:%+v %s", filter, err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	// 查询用户信息
	openids := make([]string, 0)
	for openid, _ := range usergeos {
		openids = append(openids, openid)
	}
	userinfos, err := db.GetUsers(ctx, openids)
	if err != nil {
		log.Printf("GetUsersByGeo db.GetUsers openids:%+v %s", openids, err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	resp := GetUsersByGeoResp{
		List: []*GetUsersByGeoResp_Item{},
	}
	for openid, geo := range usergeos {
		uinfo := userinfos[openid]
		feature := fmt.Sprintf("%d.%d.%d.%d", uinfo.Height, uinfo.Weight, uinfo.Age, uinfo.Length)
		item := &GetUsersByGeoResp_Item{
			Latitude:  geo.Latitude,
			Longitude: geo.Longitude,
			Avatar:    uinfo.Avatar,
			Feature:   feature,
			WeixinID:  uinfo.WeixinID,
		}
		resp.List = append(resp.List, item)
	}
	ctx.JSON(http.StatusOK, resp)
}
