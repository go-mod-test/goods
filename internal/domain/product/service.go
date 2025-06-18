package product

import (
	"context"

	"github.com/go-mod-test/goods/internal/domain"
)

type GetAllProducts interface {
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
}
type GetOneProduct interface {
	GetOneProduct(ctx context.Context, id int) (domain.Product, error)
}

type CreateProduct interface {
	CreateProduct(ctx context.Context, product domain.Product) error
}

type UpdateProduct interface {
	UpdateProduct(ctx context.Context, product domain.Product) error
}

type DeleteProduct interface {
	DelOneProduct(ctx context.Context, id int) error
}
