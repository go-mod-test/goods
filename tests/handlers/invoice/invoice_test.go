package invoice_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-mod-test/goods/internal/domain"
	"github.com/go-mod-test/goods/internal/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllInvoices(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	tim := time.Now()

	mock := &mockInvoiceGetter{
		mockData: []domain.Invoice{
			{ID: 1, Number: "123", CreatedAt: tim, CustomerID: 1},
			{ID: 2, Number: "456", CreatedAt: tim, CustomerID: 2},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/invoice/all", nil)
	w := httptest.NewRecorder()

	handler := handlers.GetAllInvoices(logger, mock)

	router := chi.NewRouter()
	router.Get("/invoice/all", handler)
	router.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var invoice []domain.Invoice
	err := json.NewDecoder(resp.Body).Decode(&invoice)
	require.NoError(t, err)

	assert.Len(t, invoice, 2)
	assert.Equal(t, "123", invoice[0].Number)
}

func TestGetOneInvoice(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	tim := time.Now()
	mock := &mockInvoiceGetter{
		oneMockData: domain.Invoice{ID: 1, Number: "123", CreatedAt: tim, CustomerID: 1},
	}

	req := httptest.NewRequest(http.MethodGet, "/invoice/1", nil)
	w := httptest.NewRecorder()

	handler := handlers.GetOneInvoice(logger, mock)

	router := chi.NewRouter()
	router.Get("/invoice/{id}", handler)
	router.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var inv domain.Invoice
	err := json.NewDecoder(resp.Body).Decode(&inv)
	require.NoError(t, err)

	assert.Equal(t, "123", inv.Number)
}

func TestCreateInvoice(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	newInvoice := domain.Invoice{Number: "123", CreatedAt: time.Now(), CustomerID: 5}
	body, err := json.Marshal(newInvoice)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/invoice/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mock := &mockInvoiceGetter{}

	handler := handlers.CreateInvoice(logger, mock)
	router := chi.NewRouter()
	router.Post("/invoice/create", handler)
	router.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestUpdateInvoice(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	newInvoice := domain.Invoice{Number: "123", CreatedAt: time.Now(), CustomerID: 5}
	body, err := json.Marshal(newInvoice)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, "/invoice/update/1", bytes.NewReader(body))

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mock := &mockInvoiceGetter{}

	handler := handlers.UpdateInvoice(logger, mock)
	router := chi.NewRouter()
	router.Put("/invoice/update/{id}", handler)
	router.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

}

func TestDeleteInvoice(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	req := httptest.NewRequest(http.MethodDelete, "/invoice/delete/1", nil)
	w := httptest.NewRecorder()

	mock := &mockInvoiceGetter{}

	handler := handlers.DeleteInvoice(logger, mock)
	router := chi.NewRouter()
	router.Delete("/invoice/delete/{id}", handler)
	router.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
