package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/cmfunc/jipeng/db"
	"github.com/gin-gonic/gin"
)

type UploadUserinfoRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Feature  string `json:"feature"`
	WeixinID string `json:"weixinID"`
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
	if param.Avatar != "" {
		update["avatar"] = param.Avatar
	}
	if param.Feature != "" {
		features := strings.Split(param.Feature, ".")
		if len(features) != 4 {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		height, err := strconv.Atoi(features[0])
		if err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		if height > 230 || height < 100 {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		weight, err := strconv.Atoi(features[1])
		if err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		if weight > 150 || weight < 45 {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		age, err := strconv.Atoi(features[2])
		if err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		if age > 80 || age < 18 {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		length, err := strconv.Atoi(features[3])
		if err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		if length > 25 || length < 2 {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		update["height"] = height
		update["weight"] = weight
		update["age"] = age
		update["length"] = length
	}
	if param.WeixinID != "" {
		update["weixin_id"] = param.WeixinID
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
