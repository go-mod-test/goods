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

func GetOneInvoiceItem(log *slog.Logger, cust invoice.GetItemInvo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.GetOneInvoiceItem"
		log.Info("get one InvoiceItem")

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

		w.Header().Set("Content-Type", "application/json")

		outprod, err := cust.GetAllProductFromInvoice(r.Context(), id)
		if err != nil {
			log.Error("failed to get InvoiceItem", "error", err, "op", op)
			http.Error(w, "InvoiceItem not found", http.StatusNotFound)
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

	}
}

func CreateInvoiceItem(log *slog.Logger, cust invoice.AddItemInvo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.CreateInvoiceItem"
		log.Info("create InvoiceItem")

		if r.Body == nil {
			log.Error("empty body", "op", op)
			http.Error(w, "empty body", http.StatusBadRequest)
			return
		}

		var temp domain.AddItemInput
		if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
			log.Error("invalid json", "error", err, "op", op)
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		if temp.InvoiceID == 0 || temp.InvoiceID < 0 || temp.ProductID == 0 || temp.ProductID < 0 || temp.Quantity < 0 {
			log.Error("empty name", "op", op)
			http.Error(w, "empty name", http.StatusBadRequest)
			return
		}

		err := cust.CreateProductInInvoice(r.Context(), temp)
		if err != nil {
			log.Error("error updating", "error", err, "op", op)
			http.Error(w, "error updating", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func DeleteInvoiceItem(log *slog.Logger, cust invoice.DelItemInvo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.DeleteInvoiceItem"
		log.Info("delete InvoiceItem")

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

		err = cust.DelProductFromInvoice(r.Context(), id)
		if err != nil {
			log.Error("error updating", "error", err, "op", op)
			http.Error(w, "error updating", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	}
}
