package handlers

type Userinfo struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Feature  string `json:"feature"`
	WeixinID string `json:"weixinID"`
}
