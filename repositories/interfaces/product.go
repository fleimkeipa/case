package interfaces

import "github.com/fleimkeipa/case/model"

type ProductRepository interface {
	FindAll(suplierID string) (*model.ProductsResponse, error)
	FindOne(id string) (*model.Product, error)
}
