package minihandlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LeaveAMessageRequest struct {
	RecvOpenid string  `json:"recv_openid,omitempty"`
	Content    string  `json:"content,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
}

// leave a message to someone, when reciver is online, reciver's map will
// show other user's message for reciver.
func LeaveAMessage(ctx *gin.Context) {
	param := LeaveAMessageRequest{}
	err := ctx.BindJSON(&param)
	if err != nil {
		log.Printf("LeaveAMessage param:%+v %s", param, err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	// cache user message into redis

}

type GetLeaveMessageListRequest struct {
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	MessageType string  `json:"message_type,omitempty"` //private, public, all（use redis union mutil set tech）
}

// 获取留言列表
func GetLeaveMessageList(ctx *gin.Context) {

}

// 获取与单个用户的留言信息
func GetLeaveMessages(ctx *gin.Context) {

}
