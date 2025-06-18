package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-mod-test/goods/internal/domain"
	"github.com/go-mod-test/goods/internal/domain/customer"
)

func GetAllCustomers(log *slog.Logger, cust customer.GetAllCustom) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.GetAllCustomers"
		log.Debug("request started", "op", op)

		w.Header().Set("Content-Type", "application/json")

		outprod, err := cust.GetAllCustomers(r.Context())
		if err != nil {
			log.Error("failed to get customers", "error", err, "op", op)
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

func GetOneCustomer(log *slog.Logger, cust customer.GetOneCustom) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.GetOneCustomer"
		log.Info("get one customer")

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

		outprod, err := cust.GetOneCustomer(r.Context(), id)
		if err != nil {
			log.Error("failed to get customer", "error", err, "op", op)
			http.Error(w, "customer not found", http.StatusNotFound)
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

func CreateCustomer(log *slog.Logger, cust customer.CreateCustom) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.CreateCustomer"
		log.Info("create customer")

		if r.Body == nil {
			log.Error("empty body", "op", op)
			http.Error(w, "empty body", http.StatusBadRequest)
			return
		}

		var req domain.Customer
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("invalid json", "error", err, "op", op)
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		if req.FirstName == "" || req.LastName == "" {
			log.Error("empty name", "op", op)
			http.Error(w, "empty name", http.StatusBadRequest)
			return
		}

		err := cust.CreateCustomer(r.Context(), req)
		if err != nil {
			log.Error("error updating", "error", err, "op", op)
			http.Error(w, "error updating", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateCustomer(log *slog.Logger, cust customer.UpdateCustom) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.UpdateCustomer"
		log.Info("update customer")
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

		var req domain.Customer
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("invalid json", "error", err, "op", op)
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		if req.FirstName == "" || req.LastName == "" {
			log.Error("empty name", "op", op)
			http.Error(w, "empty name", http.StatusBadRequest)
			return
		}

		req.ID = id

		err = cust.UpdateCustomer(r.Context(), id, req)
		if err != nil {
			log.Error("error updating", "error", err, "op", op)
			http.Error(w, "error updating", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteCustomer(log *slog.Logger, cust customer.DelOneCustom) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.DeleteCustomer"
		log.Info("delete customer")

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

		err = cust.DelOneCustomer(r.Context(), id)
		if err != nil {
			log.Error("error updating", "error", err, "op", op)
			http.Error(w, "error updating", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
