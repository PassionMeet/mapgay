package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cmfunc/jipeng/db"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Code string `json:"code"`
}

type LoginResponse struct {
	Openid string `json:"openid"`
}

type WxLoginResponse struct {
	SessionKey string `json:"session_key"` //会话密钥
	Unionid    string `json:"unionid"`     //用户在开放平台唯一标识符
	Errmsg     string `json:"errmsg"`      //错误信息
	Openid     string `json:"openid"`      //用户唯一标识
	Errcode    int32  `json:"errcode"`     //错误码
}

// Login 登陆
func Login(ctx *gin.Context) {
	param := LoginRequest{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	// wx login
	url := "https://api.weixin.qq.com/sns/jscode2session"

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	query := req.URL.Query()
	query.Add("appid", "wx9f7bd3e49f313011")
	query.Add("secret", "1674a41dbdd345548f74d8e879aafe51")
	query.Add("js_code", param.Code)
	query.Add("grant_type", "authorization_code")
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	defer resp.Body.Close()
	// 解析resp
	wxsession := WxLoginResponse{}
	err = json.NewDecoder(resp.Body).Decode(&wxsession)
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

	ctx.JSON(http.StatusOK, &LoginResponse{Openid: wxsession.Openid})

}
