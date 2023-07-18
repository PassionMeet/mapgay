package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

const RedisPrefixKey = "cos::auth::%s::%s" // cos授权 cos::auth::[situation]::[openid]

type Situation string

const SituationUploadAvatar Situation = "UploadAvatar"
const SituationUploadAvatarExpiration time.Duration = time.Hour * 2

func GetCosAuth(ctx context.Context, situation Situation, openid string) (val *sts.CredentialResult, err error) {
	key := fmt.Sprintf(RedisPrefixKey, situation, openid)
	err = redisClient.Get(ctx, key).Scan(val)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func SetCosAuth(ctx context.Context, situation Situation, openid string, val *sts.CredentialResult) (err error) {
	key := fmt.Sprintf(RedisPrefixKey, situation, openid)
	valbytes, err := json.Marshal(val)
	if err != nil {
		return err
	}
	res, err := redisClient.Set(ctx, key, valbytes, SituationUploadAvatarExpiration).Result()
	if err != nil {
		return err
	}
	log.Printf("SetCosAuth redisClient.Set %s", res)
	return nil
}
