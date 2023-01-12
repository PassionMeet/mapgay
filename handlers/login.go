package handlers

import (
	"net/http"

	"github.com/cmfunc/jipeng/db"
	"github.com/cmfunc/jipeng/wx"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Code string `json:"code"`
}

type LoginResponse struct {
	Openid string `json:"openid"`
}

// Login 登陆
func Login(ctx *gin.Context) {
	param := LoginRequest{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	wxsession, err := wx.Login(ctx, param.Code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	// TODO 发送用户的登录事件到mq，由单独服务记录用户的登录行为
	// 存储sessionKey和openid
	_, err = db.InsertUsers(ctx, &db.UsersRow{
		Openid:     wxsession.Openid,
		SessionKey: wxsession.SessionKey,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, &LoginResponse{Openid: wxsession.Openid})

}
