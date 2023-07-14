package handlers

import (
	"encoding/json"
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
	Accuracy           int32   `json:"accuracy"`  //位置精确度
	Altitude           int32   `json:"altitude"`  //高度m
	ErrMsg             string  `json:"err_msg"`
	VerticalAccuracy   int32   `json:"verticalAccuracy"`   //垂直精度m
	HorizontalAccuracy int32   `json:"horizontalAccuracy"` //水平精度
}

func UploadGeo(ctx *gin.Context) {
	param := UploadGeoRequest{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	param.Openid = ctx.GetString("openid")
	// 位置信息发送到nsq
	// nsq消费者单独处理

	body, err := json.Marshal(&param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	// 保存位置信息
	log.Printf("UploadGeo %+v", string(body))
	err = cache.AddUserGeo(ctx, &cache.UserGeo{
		Openid:    param.Openid,
		Latitude:  param.Latitude,
		Longitude: param.Longitude,
	})
	if err != nil {
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
	openid := ctx.Query("openid")
	// 通过geo信息筛选出当前地图中所有用户
	// TODO 时间筛选
	filter := &cache.GeoFilter{
		Openid: openid,
	}
	usergeos, err := cache.GetUsersByGeo(ctx, filter)
	if err != nil {
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
	// TODO 通过user_id计算当前用户与地图中用户的所有匹配值，筛出匹配度较高用户
	// TODO 考虑做后期离线计算
	// TODO 考虑用户标记自己所在的大范围，只收集在同一个大范围内的用户，并异步做匹配值计算任务
	ctx.JSON(http.StatusOK, nil)
}
