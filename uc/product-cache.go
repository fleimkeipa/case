package uc

import (
	"context"
	"fmt"

	"github.com/fleimkeipa/case/repositories/interfaces"
)

type ProductCacheUC struct {
	repo interfaces.CacheRepository
}

func NewProductCacheUC(repo interfaces.CacheRepository) *ProductCacheUC {
	return &ProductCacheUC{
		repo: repo,
	}
}

func (rc *ProductCacheUC) IsExist(ctx context.Context, brandID int, barcode string) bool {
	cacheID := ProductCacheID(brandID, barcode)
	res, err := rc.repo.Exists(ctx, cacheID)
	if err != nil {
		return false
	}

	if res > 0 {
		return true
	}

	go func(ctx context.Context, cacheID string, barcode string) {
		if err := rc.repo.Set(ctx, cacheID, barcode); err != nil {
			return
		}
	}(ctx, cacheID, barcode)

	return false
}

func ProductCacheID(brandID int, barcode string) string {
	return fmt.Sprintf("product:%d:%v", brandID, barcode)
}
