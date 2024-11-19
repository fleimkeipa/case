package interfaces

import (
	"context"

	"github.com/fleimkeipa/case/model"
)

type CacheRepository interface {
	Set(ctx context.Context, key string, value model.Product) error
	Get(ctx context.Context, key string) (*model.Product, error)
	Exists(ctx context.Context, keys ...string) bool
}
