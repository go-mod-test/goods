package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-mod-test/goods/internal/domain"
	"github.com/go-mod-test/goods/internal/domain/invoice"
)

func GetAllInvoices(log *slog.Logger, cust invoice.AllInvoiceGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.GetAllInvoices"
		log.Debug("request started", "op", op)

		w.Header().Set("Content-Type", "application/json")

		outprod, err := cust.GetAllInvoices(r.Context())
		if err != nil {
			log.Error("failed to get invoices", "error", err, "op", op)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(outprod)
		if err != nil {
			log.Error("failed serelisation", "error", err, "op", op)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		_, writeErr := w.Write(result)
		if writeErr != nil {
			log.Error("failed to write response", "error", writeErr)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
		log.Debug("request finished", "op", op)
	}
}

func GetOneInvoice(log *slog.Logger, cust invoice.OneInvoiceGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.GetOneInvoice"
		log.Info("get one invoice")

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error("invalid id")
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		if id == 0 || id < 0 {
			log.Error("empty id", "op", op)
			http.Error(w, "empty id", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		outprod, err := cust.GetOneInvoice(r.Context(), id)
		if err != nil {
			log.Error("failed to get Invoice", "error", err, "op", op)
			http.Error(w, "Invoice not found", http.StatusNotFound)
			return
		}

		result, err := json.Marshal(outprod)
		if err != nil {
			log.Error("failed serelisation", "error", err, "op", op)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		_, writeErr := w.Write(result)
		if writeErr != nil {
			log.Error("failed to write response", "error", writeErr, "op", op)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
		log.Debug("request finished", "op", op)

	}
}

func CreateInvoice(log *slog.Logger, cust invoice.InvoiceCreater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.CreateInvoice"
		log.Info("create Invoice")

		if r.Body == nil {
			log.Error("empty body", "op", op)
			http.Error(w, "empty body", http.StatusBadRequest)
			return
		}

		var req domain.Invoice
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("invalid json", "error", err, "op", op)
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		if req.CustomerID == 0 || req.CustomerID < 0 || req.Number == "" {
			log.Error("empty customer id", "op", op)
			http.Error(w, "empty customer id", http.StatusBadRequest)
			return
		}

		err := cust.CreateInvoice(r.Context(), req)
		if err != nil {
			log.Error("error updating", "error", err, "op", op)
			http.Error(w, "error updating", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateInvoice(log *slog.Logger, cust invoice.InvoiceUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.UpdateInvoice"
		log.Info("update Invoice")

		if r.Body == nil {
			log.Error("empty body", "op", op)
			http.Error(w, "empty body", http.StatusBadRequest)
			return
		}

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error("invalid id", "error", err, "op", op)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		if id == 0 || id < 0 {
			log.Error("empty id", "op", op)
			http.Error(w, "empty id", http.StatusBadRequest)
			return
		}

		var req domain.Invoice
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("invalid json", "error", err, "op", op)
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		if req.CustomerID == 0 || req.CustomerID < 0 || req.Number == "" {
			log.Error("empty customer id", "op", op)
			http.Error(w, "empty customer id", http.StatusBadRequest)
			return
		}

		req.ID = id

		err = cust.UpdateInvoice(r.Context(), req)
		if err != nil {
			log.Error("error updating", "error", err, "op", op)
			http.Error(w, "error updating", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteInvoice(log *slog.Logger, cust invoice.OneInvoiceDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.DeleteInvoice"
		log.Info("delete Invoice")

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error("invalid id", "error", err, "op", op)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		if id == 0 || id < 0 {
			log.Error("empty id", "op", op)
			http.Error(w, "empty id", http.StatusBadRequest)
			return
		}

		err = cust.DelOneInvoice(r.Context(), id)
		if err != nil {
			log.Error("error updating", "error", err, "op", op)
			http.Error(w, "error updating", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
