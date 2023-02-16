package wx

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/cmfunc/jipeng/conf"
)

const (
	url = "https://api.weixin.qq.com/sns/jscode2session"
)

var _wx = &conf.Wx{}

func Set(cfg *conf.Wx) {
	_wx = cfg
}

func Get() *conf.Wx {
	return _wx
}

type WxLoginResponse struct {
	SessionKey string `json:"session_key"` //会话密钥
	Unionid    string `json:"unionid"`     //用户在开放平台唯一标识符
	Errmsg     string `json:"errmsg"`      //错误信息
	Openid     string `json:"openid"`      //用户唯一标识
	Errcode    int32  `json:"errcode"`     //错误码
}

func Login(ctx context.Context, _conf conf.Wx, code string) (*WxLoginResponse, error) {
	// wx login

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("appid", _conf.AppID)
	query.Add("secret", _conf.AppSecret)
	query.Add("js_code", code)
	query.Add("grant_type", "authorization_code")
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	// 解析resp
	wxsession := WxLoginResponse{}
	err = json.NewDecoder(resp.Body).Decode(&wxsession)
	if err != nil {
		return nil, err
	}
	return &wxsession, nil
}
