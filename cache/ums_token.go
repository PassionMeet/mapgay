package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

const (
	KeyPrefixUmsToken = "ums::token::%s" //ums::token::[account]
)

const (
	KeyTTLUmsToken = time.Hour * 24 * 30
)

func SetUMSToken(ctx context.Context, account, token string) error {
	key := fmt.Sprintf(KeyPrefixUmsToken, account)
	_, err := redisClient.Set(ctx, key, token, KeyTTLUmsToken).Result()
	if err != nil {
		return errors.Wrapf(err, "SetUMSToken %s %s ", account, token)
	}
	return nil
}
func GetUMSToken(ctx context.Context, account string) (token string, err error) {
	key := fmt.Sprintf(KeyPrefixUmsToken, account)
	token, err = redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", errors.Wrapf(err, "SetUMSToken %s %s ", account, token)
	}
	return token, nil
}
