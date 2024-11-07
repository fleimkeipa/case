package uc

import (
	"github.com/fleimkeipa/case/model"
	"github.com/fleimkeipa/case/repositories/interfaces"
)

type ProductUC struct {
	repo interfaces.ProductRepository
}

func NewProductUC(repo interfaces.ProductRepository) *ProductUC {
	return &ProductUC{repo: repo}
}

func (rc *ProductUC) FindAll(suplierID string) (*model.ProductsResponse, error) {
	return rc.repo.FindAll(suplierID)
}

func (rc *ProductUC) FindOne(id string) (*model.Product, error) {
	return rc.repo.FindOne(id)
}
