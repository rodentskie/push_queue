package redis

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/redis/go-redis/v9"
)

func New(addr, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: db})
}

func PutID(ctx context.Context, rdb *redis.Client, sheet, originalID string, generated []byte, ttl time.Duration) error {
	key := sheet + ":" + originalID
	val := base64.StdEncoding.EncodeToString(generated)
	return rdb.Set(ctx, key, val, ttl).Err()
}

func GetID(ctx context.Context, rdb *redis.Client, sheet, originalID string) ([]byte, error) {
	key := sheet + ":" + originalID
	s, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(s)
}
