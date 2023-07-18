package storage

// 腾讯对象存储
import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cmfunc/jipeng/conf"
	"github.com/tencentyun/cos-go-sdk-v5"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

const AvatarCosRegion = "https://avatar-1256498508.cos.ap-beijing.myqcloud.com"

var client *cos.Client

func InitCos(config *conf.Cos) {
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(config.Url)
	b := &cos.BaseURL{BucketURL: u}
	client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: config.SecretID, // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: config.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

}

func GetCosStsCredential(config *conf.Cos, cosRegion, openid string) (*sts.Credentials, error) {
	c := sts.NewClient(config.SecretID, config.SecretKey, nil)
	// 策略概述 https://cloud.tencent.com/document/product/436/18023
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          "ap-guangzhou",
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
					},
					Effect: "allow",
					Resource: []string{
						//这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						//存储桶的命名格式为 BucketName-APPID，此处填写的 bucket 必须为此格式
						// "qcs::cos:" + region + ":uid/" + appid + ":" + bucket + "/exampleobject",
						cosRegion,
					},
					// 开始构建生效条件 condition
					// 关于 condition 的详细设置规则和COS支持的condition类型可以参考https://cloud.tencent.com/document/product/436/71306
					Condition: map[string]map[string]interface{}{
						// "ip_equal": map[string]interface{}{
						// 	"qcs:ip": []string{
						// 		"10.217.182.3/24",
						// 		"111.21.33.72/24",
						// 	},
						// },
					},
				},
			},
		},
	}
	res, err := c.GetCredential(opt)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", res.Credentials)
	return res.Credentials, nil
}
