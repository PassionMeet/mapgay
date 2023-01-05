package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cmfunc/jipeng/mq"
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
	// 位置信息发送到nsq
	// nsq消费者单独处理

	body, err := json.Marshal(&param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	err = mq.PubUserGeo(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, nil)

}

type GetUsersByGeoParam struct {
	Openid    string  `form:"openid"`
	Latitude  float64 `form:"latitude"`
	Longitude float64 `form:"longitude"`
}

type GetUsersByGeoResp_Item struct {
	// 和微信小程序map组建里的marker保持相似
}

type GetUsersByGeoResp struct {
	List []*GetUsersByGeoResp_Item `json:"list"`
}

func GetUsersByGeo(ctx *gin.Context) {
	// 通过geo信息筛选出当前地图中所有用户

	// 通过user_id计算当前用户与地图中用户的所有匹配值，筛出匹配度较高用户

	// TODO 考虑做后期离线计算
	// TODO 考虑用户标记自己所在的大范围，只收集在同一个大范围内的用户，并异步做匹配值计算任务
}
