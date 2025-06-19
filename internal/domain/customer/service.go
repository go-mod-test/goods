package customer

import (
	"context"

	"github.com/go-mod-test/goods/internal/domain"
)

type AllCustomerGetter interface {
	GetAllCustomers(ctx context.Context) ([]domain.Customer, error)
}

type OneCustomerGetter interface {
	GetOneCustomer(ctx context.Context, id int) (domain.Customer, error)
}

type CustomerCreater interface {
	CreateCustomer(ctx context.Context, customer domain.Customer) error
}

type CustomerUpdater interface {
	UpdateCustomer(ctx context.Context, customer domain.Customer) error
}

type OneCustomerDeleter interface {
	DelOneCustomer(ctx context.Context, id int) error
}
