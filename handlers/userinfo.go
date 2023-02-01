package handlers

import (
	"net/http"

	"github.com/cmfunc/jipeng/db"
	"github.com/gin-gonic/gin"
)

type UploadUserinfoRequest struct {
	Nickname string `json:"nickname"`
	Avator   string `json:"avator"`
}
type UploadUserinfoResponse struct{}

// UploadUserinfo 上传用户个人信息
func UploadUserinfo(ctx *gin.Context) {
	openid := ctx.GetString("openid")
	param := UploadUserinfoRequest{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	where := map[string]interface{}{"openid": openid}
	update := map[string]interface{}{}
	if param.Nickname != "" {
		update["username"] = param.Nickname
	}
	if param.Avator != "" {
		update["avator"] = param.Avator
	}
	if len(update) < 1 {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	err = db.UpdateUser(ctx, where, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
