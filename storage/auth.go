package storage

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func GetCosBucketHost(bucket, appid, region string) string {
	return fmt.Sprintf("https://%s-%s.cos.%s.myqcloud.com", bucket, appid, region)
}

func GetCosAuth(bucket, object, appid, region, token, secretid, secretkey string) (requestUrl *url.URL, err error) {
	host := GetCosBucketHost(bucket, appid, region)
	u, _ := url.Parse(host)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{})

	ctx := context.Background()

	// 方法1 通过 PresignedURLOptions 设置 x-cos-security-token
	// PresignedURLOptions 提供用户添加请求参数和请求头部
	opt := &cos.PresignedURLOptions{
		Query:  &url.Values{},
		Header: &http.Header{},
	}
	opt.Query.Add("x-cos-security-token", token)
	// 获取预签名
	presignedURL, err := c.Object.GetPresignedURL(ctx, http.MethodPut, object, secretid, secretkey, time.Hour, opt)
	if err != nil {
		return nil, err
	}
	return presignedURL, nil
}
