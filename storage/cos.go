package storage

// 腾讯对象存储
import (
	"fmt"
	"time"

	"github.com/cmfunc/jipeng/conf"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

func GetCosStsCredential(config conf.Cos, bucketName, appid, region, openid string) (*sts.CredentialResult, error) {
	c := sts.NewClient(config.SecretID, config.SecretKey, nil)
	// 策略概述 https://cloud.tencent.com/document/product/436/18023
	sourceUrl := fmt.Sprintf("https://%s-%s.cos.%s.myqcloud.com/%s.jpg", bucketName, appid, region, openid)
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          "ap-beijing",
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
						sourceUrl,
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
	return res, nil
}
