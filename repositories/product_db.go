package repositories

import (
	"fmt"

	"github.com/fleimkeipa/case/model"
	"github.com/go-pg/pg"
)

type ProductDBRepository struct {
	db *pg.DB
}

func NewProductDBRepository(db *pg.DB) *ProductDBRepository {
	return &ProductDBRepository{
		db: db,
	}
}

func (rc *ProductDBRepository) Create(product *model.Product) (*model.Product, error) {
	q := rc.db.Model(product)

	_, err := q.Insert()
	if err != nil {
		return nil, fmt.Errorf("failed to create product [%d]: %w", product.BrandID, err)
	}

	return product, nil
}
