package interfaces

import "github.com/fleimkeipa/case/model"

type ProductAPIRepository interface {
	FindAll(opts model.ProductListOpts) (*model.ProductsResponse, error)
	FindOne(id string) (*model.Product, error)
}
