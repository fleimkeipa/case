package uc

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fleimkeipa/case/model"
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

func (rc *ProductCacheUC) IsExist(ctx context.Context, suplierID, productMainID string) bool {
	cacheID := ProductCacheID(suplierID, productMainID)

	return rc.repo.Exists(ctx, cacheID)
}

func (rc *ProductCacheUC) Set(ctx context.Context, product *model.Product) error {
	cacheID := ProductCacheID(strconv.Itoa(product.SupplierID), product.ProductMainID)

	if err := rc.repo.Set(ctx, cacheID, *product); err != nil {
		return err
	}

	return nil
}

func (rc *ProductCacheUC) Get(ctx context.Context, suplierID, productMainID string) (*model.Product, error) {
	cacheID := ProductCacheID(suplierID, productMainID)

	product, err := rc.repo.Get(ctx, cacheID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func ProductCacheID(suplierID, productMainID string) string {
	return fmt.Sprintf("product:%v:%v", suplierID, productMainID)
}
