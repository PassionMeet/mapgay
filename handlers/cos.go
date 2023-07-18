package handlers

import (
	"log"
	"net/http"

	"github.com/cmfunc/jipeng/cache"
	"github.com/cmfunc/jipeng/conf"
	"github.com/cmfunc/jipeng/storage"
	"github.com/gin-gonic/gin"
)

type GetCosAuthRequest struct {
	Bucket string `form:"bucket,omitempty"`
	Region string `form:"region,omitempty"`
}

// GetCosAuth 获取cos授权
func GetCosAuth(ctx *gin.Context) {
	param := GetCosAuthRequest{}
	err := ctx.BindJSON(&param)
	if err != nil {
		log.Printf("GetCosAuth param:%+v %s", param, err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	openid := ctx.GetString("openid")
	// 获取redis中，cos授权
	authVal, err := cache.HGetCosAuth(ctx, cache.SituationUploadAvatar)
	if err == nil && authVal != nil {
		ctx.JSON(http.StatusOK, authVal)
		return
	}
	log.Printf("GetCosAuth cache.HGetCosAuth situation:%s %s", cache.SituationUploadAvatar, err)
	// redis中未获取到，使用Cos工具进行授权
	cosRegion := storage.AvatarCosRegion + "/" + "openid" + ".jpg"
	authVal, err = storage.GetCosStsCredential(conf.Get().Cos, cosRegion, openid)
	if err != nil {
		log.Printf("GetCosAuth storage.GetCosStsCredential %+v %s %s", conf.Get().Cos, cosRegion, err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, authVal)
}
