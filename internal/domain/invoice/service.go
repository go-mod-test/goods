package invoice

import (
	"context"

	"github.com/go-mod-test/goods/internal/domain"
)

type GetOneInvo interface {
	GetOneInvoice(ctx context.Context, id int) (domain.Invoice, error)
}

type GetAllInvo interface {
	GetAllInvoices(ctx context.Context) ([]domain.Invoice, error)
}

type CreateInvo interface {
	CreateInvoice(ctx context.Context, invoice domain.Invoice) error
}

type UpdateInvo interface {
	UpdateInvoice(ctx context.Context, invoice domain.Invoice) error
}

type DelOneInvo interface {
	DelOneInvoice(ctx context.Context, id int) error
}

type AddItemInvo interface {
	CreateProductInInvoice(ctx context.Context, input domain.AddItemInput) error
}

type DelItemInvo interface {
	DelProductFromInvoice(ctx context.Context, invoiceId int) error
}

type GetItemInvo interface {
	GetAllProductFromInvoice(ctx context.Context, invoiceId int) ([]domain.InvoiceItem, error)
}
