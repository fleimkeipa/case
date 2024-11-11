package repositories

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	Client *redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{
		Client: client,
	}
}

func (rc *CacheRepository) Set(ctx context.Context, key string, value string) error {
	return rc.Client.Set(ctx, key, value, 0).Err()
}

func (rc *CacheRepository) Get(ctx context.Context, key string) (string, error) {
	return rc.Client.Get(ctx, key).Result()
}

func (rc *CacheRepository) Exists(ctx context.Context, keys ...string) (int64, error) {
	return rc.Client.Exists(ctx, keys...).Result()
}
