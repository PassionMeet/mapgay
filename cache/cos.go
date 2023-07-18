package cache

import (
	"context"
	"fmt"

	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

const RedisPrefixKey = "cos::auth::%s" // cos授权 cos::auth::[situation]

type CosAuth struct {
}

type Situation string

const SituationUploadAvatar Situation = "UploadAvatar"

func HGetCosAuth(ctx context.Context, situation Situation) (val *sts.Credentials, err error) {
	key := fmt.Sprintf(RedisPrefixKey, situation)
	err = redisClient.HGetAll(ctx, key).Scan(val)
	if err != nil {
		return nil, err
	}
	return val, nil
}
