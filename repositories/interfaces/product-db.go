package interfaces

import "github.com/fleimkeipa/case/model"

type ProductDBRepository interface {
	Create(product *model.Product) (*model.Product, error)
}
