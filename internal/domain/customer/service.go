package customer

import (
	"context"

	"github.com/go-mod-test/goods/internal/domain"
)

type GetAllCustom interface {
	GetAllCustomers(ctx context.Context) ([]domain.Customer, error)
}

type GetOneCustom interface {
	GetOneCustomer(ctx context.Context, id int) (domain.Customer, error)
}

type CreateCustom interface {
	CreateCustomer(ctx context.Context, customer domain.Customer) error
}

type UpdateCustom interface {
	UpdateCustomer(ctx context.Context, id int, customer domain.Customer) error
}

type DelOneCustom interface {
	DelOneCustomer(ctx context.Context, id int) error
}
