package handlers

import (
	"fmt"
	"net/http"

	"github.com/cmfunc/jipeng/cache"
	"github.com/cmfunc/jipeng/conf"
	"github.com/cmfunc/jipeng/db"
	"github.com/cmfunc/jipeng/wx"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Code string `json:"code"`
}

type LoginResponse struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"sessionKey"`
	Userinfo
}

// Login 登陆
func Login(ctx *gin.Context) {
	param := LoginRequest{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	wxsession, err := wx.Login(ctx, *conf.Get().Wx, param.Code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	// 存储sessionKey和openid
	_, err = db.InsertUsers(ctx, &db.UsersRow{
		Openid:     wxsession.Openid,
		SessionKey: wxsession.SessionKey,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	// 缓存sessionKey和openid
	_, err = cache.SetUserSession(ctx, wxsession.Openid, wxsession.SessionKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	// 查询用户信息
	userRow, err := db.GetUser(ctx, wxsession.Openid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	// 登陆成功以后，前端接受到openid和sessionkey需要在请求时写入header中，并由中间件获取校验
	resp := &LoginResponse{
		Openid:     wxsession.Openid,
		SessionKey: wxsession.SessionKey,
		Userinfo: Userinfo{
			Nickname: userRow.Username,
			Avatar:   userRow.Avatar,
			WeixinID: userRow.WeixinID,
		},
	}
	if userRow.Height > 0 && userRow.Weight > 0 && userRow.Age > 0 && userRow.Length > 0 {
		resp.Feature = fmt.Sprintf("%d.%d.%d.%d", userRow.Height, userRow.Weight, userRow.Age, userRow.Length)
	}
	ctx.JSON(http.StatusOK, resp)

}
