package customer_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-mod-test/goods/internal/domain"
	"github.com/go-mod-test/goods/internal/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllCustomers(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil)) // глушим лог в тестах

	mock := &mockCustomerGetter{
		mockData: []domain.Customer{
			{ID: 1, FirstName: "Иван", LastName: "Петров"},
			{ID: 2, FirstName: "Анна", LastName: "Сидорова"},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/customer/all", nil)
	w := httptest.NewRecorder()

	handler := handlers.GetAllCustomers(logger, mock)
	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var customers []domain.Customer
	err := json.NewDecoder(resp.Body).Decode(&customers)
	require.NoError(t, err)

	assert.Len(t, customers, 2)
	assert.Equal(t, "Иван", customers[0].FirstName)
}

func TestGetOneCustomer(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	mock := &mockCustomerGetter{
		oneMockData: domain.Customer{ID: 1, FirstName: "Иван", LastName: "Петров"},
	}

	req := httptest.NewRequest(http.MethodGet, "/customer/1", nil)
	w := httptest.NewRecorder()

	handler := handlers.GetOneCustomer(logger, mock)
	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var customer domain.Customer
	err := json.NewDecoder(resp.Body).Decode(&customer)
	require.NoError(t, err)

	assert.Equal(t, "Иван", customer.FirstName)
}

func TestCreateCustomer(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	newCustomer := domain.Customer{FirstName: "Мария", LastName: "Иванова"}
	body, err := json.Marshal(newCustomer)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/customer/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mock := &mockCustomerGetter{}

	handler := handlers.CreateCustomer(logger, mock)
	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestUpdateCustomer(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	newCustomer := domain.Customer{ID: 1, FirstName: "Мария", LastName: "Иванова"}
	body, err := json.Marshal(newCustomer)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, "/customer/update/1", bytes.NewReader(body))

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mock := &mockCustomerGetter{}

	handler := handlers.UpdateCustomer(logger, mock)
	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

}

func TestDeleteCustomer(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	req := httptest.NewRequest(http.MethodDelete, "/customer/delete/1", nil)
	w := httptest.NewRecorder()

	mock := &mockCustomerGetter{}

	handler := handlers.DeleteCustomer(logger, mock)
	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
