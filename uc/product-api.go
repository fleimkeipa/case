package uc

import (
	"context"
	"log"

	"github.com/fleimkeipa/case/model"
	"github.com/fleimkeipa/case/repositories/interfaces"
)

type ProductAPIUC struct {
	api interfaces.ProductAPIRepository
	db  ProductDBUC
}

func NewProductAPIUC(api interfaces.ProductAPIRepository, db ProductDBUC) *ProductAPIUC {
	return &ProductAPIUC{
		api: api,
		db:  db,
	}
}

func (rc *ProductAPIUC) FindAll(ctx context.Context, suplierID string) (*model.ProductsResponse, error) {
	res, err := rc.api.FindAll(suplierID)
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

func (rc *ProductAPIUC) FindOne(id string) (*model.Product, error) {
	return rc.api.FindOne(id)
}
