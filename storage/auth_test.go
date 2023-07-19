package storage

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/cmfunc/jipeng/conf"
	"github.com/spf13/viper"
)

func TestGetCosAuth(t *testing.T) {
	viper.SetConfigFile("../config.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var _config = conf.Config{}

	err = viper.Unmarshal(&_config)
	if err != nil {
		panic(err)
	}
	openid := "oxoKy5JA3e2JIQzJU09gR9-MdKVA"
	extFile := "jpg"
	credToken, err := GetCosStsCredential(_config.Cos, _config.Cos.Avatar.BucketName, _config.Cos.Avatar.Appid, _config.Cos.Avatar.Region, openid, extFile)
	if err != nil {
		panic(err)
	}
	objectname := openid + "." + extFile
	url, err := GetCosAuth(_config.Cos.Avatar.BucketName,
		objectname,
		_config.Cos.Avatar.Appid,
		_config.Cos.Avatar.Region,
		credToken.Credentials.SessionToken,
		credToken.Credentials.TmpSecretID,
		credToken.Credentials.TmpSecretKey)
	if err != nil {
		panic(err)
	}

	// 通过预签名方式上传对象
	data := "test upload with presignedURL"
	f := strings.NewReader(data)
	req, err := http.NewRequest(http.MethodPut, url.String(), f)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	t.Logf("%+v", url)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	t.Logf("%+v", string(body))

}
