package interfaces

import "github.com/fleimkeipa/case/model"

type ProductAPIRepository interface {
	FindAll(opts model.ProductListOpts) (*model.ProductsResponse, error)
	FindOne(suplierID, productMainID string) (*model.Product, error)
}
