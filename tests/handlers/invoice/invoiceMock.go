package invoice_test

import (
	"context"

	"github.com/go-mod-test/goods/internal/domain"
)

type mockInvoiceGetter struct {
	mockData    []domain.Invoice
	oneMockData domain.Invoice
	err         error
}

func (m *mockInvoiceGetter) GetAllInvoices(_ context.Context) ([]domain.Invoice, error) {
	return m.mockData, m.err
}

func (m *mockInvoiceGetter) GetOneInvoice(_ context.Context, _ int) (domain.Invoice, error) {
	return m.oneMockData, m.err
}

func (m *mockInvoiceGetter) CreateInvoice(_ context.Context, _ domain.Invoice) error {
	return m.err
}

func (m *mockInvoiceGetter) UpdateInvoice(_ context.Context, _ domain.Invoice) error {
	return m.err
}

func (m *mockInvoiceGetter) DelOneInvoice(_ context.Context, _ int) error {
	return m.err
}
