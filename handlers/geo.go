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

// 抽象http接口逻辑处理
type Handle interface {
	Param(ctx *gin.Context) (interface{}, error)
	Service(interface {
		User()
	})
	Resp(interface{})
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
