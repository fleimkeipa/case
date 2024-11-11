package uc

import (
	"context"

	"github.com/fleimkeipa/case/pkg"

	"github.com/redis/go-redis/v9"
)

type ProductCacheUC struct {
	cache redis.Client
}

func NewProductCacheUC(cache redis.Client) *ProductCacheUC {
	return &ProductCacheUC{
		cache: cache,
	}
}

func (rc *ProductCacheUC) IsExist(ctx context.Context, brandID int, barcode string) bool {
	cacheID := pkg.ProductCacheID(brandID, barcode)
	res, err := rc.cache.Get(ctx, cacheID).Result()
	if err != nil {
		return false
	}

	if res != "" {
		return true
	}

	go func() {
		if err := rc.cache.Set(ctx, cacheID, barcode, 0).Err(); err != nil {
			return
		}
	}()

	return false
}
