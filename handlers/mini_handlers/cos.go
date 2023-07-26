package minihandlers

import (
	"log"
	"net/http"

	"github.com/cmfunc/jipeng/cache"
	"github.com/cmfunc/jipeng/conf"
	"github.com/cmfunc/jipeng/storage"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

type GetCosAuthRequest struct {
	Situation string `form:"situation,omitempty"`
	FileExt   string `form:"fileExt,omitempty"`
}

type GetCosAuthResponse struct {
	sts.CredentialResult
	CosHost       string
	CosKey        string
	Authorization string
	SecurityToken string
}

// GetCosAuth 获取cos授权
func GetCosAuth(ctx *gin.Context) {
	param := GetCosAuthRequest{}
	err := ctx.Bind(&param)
	if err != nil {
		log.Printf("GetCosAuth param:%+v %s", param, err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	openid := ctx.GetString("openid")
	// 获取redis中，cos授权
	authVal, err := cache.GetCosAuth(ctx, cache.SituationUploadAvatar, openid)
	if err != nil && err != redis.Nil {
		log.Printf("GetCosAuth cache.GetCosAuth %s %s", openid, err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	log.Printf("GetCosAuth cache.HGetCosAuth situation:%s %+v", cache.SituationUploadAvatar, authVal)
	if authVal == nil {
		// redis中未获取到，使用Cos工具进行授权
		avatarCosConf := conf.Get().Cos.Avatar
		authVal, err = storage.GetCosStsCredential(
			conf.Get().Cos,
			avatarCosConf.BucketName,
			avatarCosConf.Appid,
			avatarCosConf.Region,
			openid,
			param.FileExt)
		if err != nil {
			log.Printf("GetCosAuth storage.GetCosStsCredential %+v %s", conf.Get().Cos, err)
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		// 设置缓存
		err = cache.SetCosAuth(ctx, cache.SituationUploadAvatar, openid, authVal)
		if err != nil {
			log.Printf("GetCosAuth storage.SetCosAuth %+v %s %s %s", authVal, cache.SituationUploadAvatar, openid, err)
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
	}

	objectname := openid + "." + param.FileExt
	url, err := storage.GetCosAuth(conf.Get().Cos.Avatar.BucketName,
		objectname,
		conf.Get().Cos.Avatar.Appid,
		conf.Get().Cos.Avatar.Region,
		authVal.Credentials.SessionToken,
		authVal.Credentials.TmpSecretID,
		authVal.Credentials.TmpSecretKey)
	if err != nil {
		log.Printf("GetCosAuth  storage.GetCosAuth %s %+v %s", objectname, authVal, err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, url.String())
}
