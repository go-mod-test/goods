package invoice

import (
	"context"

	"github.com/go-mod-test/goods/internal/domain"
)

type OneInvoiceGetter interface {
	GetOneInvoice(ctx context.Context, id int) (domain.Invoice, error)
}

type AllInvoiceGetter interface {
	GetAllInvoices(ctx context.Context) ([]domain.Invoice, error)
}

type InvoiceCreater interface {
	CreateInvoice(ctx context.Context, invoice domain.Invoice) error
}

type InvoiceUpdater interface {
	UpdateInvoice(ctx context.Context, invoice domain.Invoice) error
}

type OneInvoiceDeleter interface {
	DelOneInvoice(ctx context.Context, id int) error
}

type ItemInvoiceAdder interface {
	CreateProductInInvoice(ctx context.Context, input domain.AddItemInput) error
}

type ItemInvoiceDeleter interface {
	DelProductFromInvoice(ctx context.Context, invoiceId int) error
}

type ItemInvoiceGetter interface {
	GetAllProductFromInvoice(ctx context.Context, invoiceId int) ([]domain.InvoiceItem, error)
}
