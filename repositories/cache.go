package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/fleimkeipa/case/model"

	"github.com/redis/go-redis/v9"
)

const expireTime = 24 * time.Hour

type CacheRepository struct {
	Client *redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{
		Client: client,
	}
}

func (rc *CacheRepository) Set(ctx context.Context, key string, value model.Product) error {
	marshalled, err := json.Marshal(&value)
	if err != nil {
		return err
	}

	return rc.Client.Set(ctx, key, string(marshalled), expireTime).Err()
}

func (rc *CacheRepository) Get(ctx context.Context, key string) (*model.Product, error) {
	var product model.Product

	res, err := rc.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	b := []byte(res)

	if err := json.Unmarshal(b, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (rc *CacheRepository) Exists(ctx context.Context, keys ...string) bool {
	count, err := rc.Client.Exists(ctx, keys...).Result()
	if err != nil {
		return false
	}

	if count > 0 {
		return true
	}

	return false
}
