package uc

import (
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

func (rc *ProductAPIUC) FindAll(suplierID string) (*model.ProductsResponse, error) {
	res, err := rc.api.FindAll(suplierID)
	if err != nil {
		return nil, err
	}

	go func(products *model.ProductsResponse) {
		for _, product := range products.Content {
			_, err := rc.db.Create(&product)
			if err != nil {
				panic(err)
			}
		}
	}(res)

	return res, nil
}

func (rc *ProductAPIUC) FindOne(id string) (*model.Product, error) {
	return rc.api.FindOne(id)
}
