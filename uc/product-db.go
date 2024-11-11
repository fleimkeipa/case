package uc

import (
	"github.com/fleimkeipa/case/model"
	"github.com/fleimkeipa/case/repositories/interfaces"
)

type ProductDBUC struct {
	repo interfaces.ProductDBRepository
}

func NewProductDBUC(repo interfaces.ProductDBRepository) *ProductDBUC {
	return &ProductDBUC{repo: repo}
}

func (rc *ProductDBUC) Create(product *model.Product) (*model.Product, error) {
	return rc.repo.Create(product)
}
