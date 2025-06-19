package invoiceItem_test

import (
	"context"

	"github.com/go-mod-test/goods/internal/domain"
)

type mockCustomerGetter struct {
	mockData    []domain.Customer
	oneMockData domain.Customer
	err         error
}

func (m *mockCustomerGetter) GetAllCustomers(_ context.Context) ([]domain.Customer, error) {
	return m.mockData, m.err
}

func (m *mockCustomerGetter) GetOneCustomer(_ context.Context, _ int) (domain.Customer, error) {
	return m.oneMockData, m.err
}

func (m *mockCustomerGetter) CreateCustomer(_ context.Context, _ domain.Customer) error {
	return m.err
}

func (m *mockCustomerGetter) UpdateCustomer(_ context.Context, _ domain.Customer) error {
	return m.err
}

func (m *mockCustomerGetter) DelOneCustomer(_ context.Context, _ int) error {
	return m.err
}
