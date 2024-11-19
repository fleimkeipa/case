package uc

import (
	"context"
	"log"

	"github.com/fleimkeipa/case/model"
	"github.com/fleimkeipa/case/repositories/interfaces"
)

type ProductAPIUC struct {
	api   interfaces.ProductAPIRepository
	db    ProductDBUC
	cache ProductCacheUC
}

func NewProductAPIUC(api interfaces.ProductAPIRepository, db ProductDBUC, cache ProductCacheUC) *ProductAPIUC {
	return &ProductAPIUC{
		api:   api,
		db:    db,
		cache: cache,
	}
}

func (rc *ProductAPIUC) FindAll(ctx context.Context, opts model.ProductListOpts) (*model.ProductsResponse, error) {
	res, err := rc.api.FindAll(opts)
	if err != nil {
		return nil, err
	}

	go func(ctx context.Context, products *model.ProductsResponse) {
		ctx = context.WithoutCancel(ctx)
		for _, product := range products.Content {
			_, err := rc.db.Create(ctx, &product)
			if err != nil {
				log.Fatalf("Failed to create product: %v", err)
			}
		}
	}(ctx, res)

	return res, nil
}

func (rc *ProductAPIUC) FindOne(ctx context.Context, suplierID, productMainID string) (*model.Product, error) {
	product, err := rc.cache.Get(ctx, suplierID, productMainID)
	if err != nil {
		return nil, err
	}

	if product != nil && product.ProductMainID == productMainID {
		return product, nil
	}

	product, err = rc.api.FindOne(suplierID, productMainID)
	if err != nil {
		return nil, err
	}

	go func(ctx context.Context, product *model.Product) {
		ctx = context.WithoutCancel(ctx)
		if err := rc.cache.Set(ctx, product); err != nil {
			return
		}
	}(ctx, product)

	return product, nil
}
