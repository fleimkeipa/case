package uc

import (
	"context"

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
	if rc.cache.IsExist(ctx, product.BrandID, product.Barcode) {
		return product, nil
	}

	return rc.repo.Create(product)
}
