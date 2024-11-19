package repositories

import (
	"fmt"
	"net/http"

	"github.com/fleimkeipa/case/model"
	"github.com/fleimkeipa/case/pkg"
)

type ProductAPIRepository struct {
	Client pkg.Client
}

func NewProductAPIRepository(client pkg.Client) *ProductAPIRepository {
	return &ProductAPIRepository{
		Client: client,
	}
}

func (rc *ProductAPIRepository) FindAll(opts model.ProductListOpts) (*model.ProductsResponse, error) {
	request := model.InternalRequest{
		Method:     http.MethodGet,
		Paths:      []string{"suppliers", opts.SuplierID.Value, "products"},
		Headers:    map[string]string{"Content-Type": "application/json"},
		Pagination: opts.PaginationOpts,
	}

	var products model.ProductsResponse
	if err := rc.Client.Do(request, &products); err != nil {
		return nil, err
	}

	return &products, nil
}

func (rc *ProductAPIRepository) FindOne(suplierID, productMainID string) (*model.Product, error) {
	request := model.InternalRequest{
		Method:  http.MethodGet,
		Paths:   []string{"suppliers", suplierID, "products"},
		Headers: map[string]string{"Content-Type": "application/json"},
		QueryParams: map[string]string{
			"productMainId": productMainID,
		},
	}

	var products model.ProductsResponse
	if err := rc.Client.Do(request, &products); err != nil {
		return nil, err
	}

	switch products.TotalElements {
	case 0:
		return nil, fmt.Errorf("product not found: [%s]", productMainID)
	case 1:
		if len(products.Content) == 0 {
			return nil, fmt.Errorf("product not found: [%s]", productMainID)
		}

		return &products.Content[0], nil
	default:
		return nil, fmt.Errorf("multiple products found: [%s]", productMainID)
	}
}
