package product

import (
	"context"

	"github.com/go-mod-test/goods/internal/domain"
)

type AllProductsGetter interface {
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
}
type OneProductGetter interface {
	GetOneProduct(ctx context.Context, id int) (domain.Product, error)
}

type ProductCreater interface {
	CreateProduct(ctx context.Context, product domain.Product) error
}

type ProductUpdater interface {
	UpdateProduct(ctx context.Context, product domain.Product) error
}

type ProductDeleter interface {
	DelOneProduct(ctx context.Context, id int) error
}
