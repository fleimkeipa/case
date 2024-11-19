package uc

import (
	"context"
	"strconv"

	"github.com/fleimkeipa/case/model"
	"github.com/fleimkeipa/case/repositories/interfaces"
)

type ProductDBUC struct {
	repo  interfaces.ProductDBRepository
	cache *ProductCacheUC
}

func NewProductDBUC(repo interfaces.ProductDBRepository, cache *ProductCacheUC) *ProductDBUC {
	return &ProductDBUC{
		repo:  repo,
		cache: cache,
	}
}

func (rc *ProductDBUC) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	if rc.cache.IsExist(ctx, strconv.Itoa(product.SupplierID), product.ProductMainID) {
		return product, nil
	}

	go func(ctx context.Context, product *model.Product) {
		if err := rc.cache.Set(ctx, product); err != nil {
			return
		}
	}(ctx, product)

	return rc.repo.Create(product)
}
