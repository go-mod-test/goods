package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-mod-test/goods/internal/domain"
	"github.com/go-mod-test/goods/internal/domain/product"
)

func GetAllProducts(log *slog.Logger, prod product.AllProductsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.GetAllProducts"

		w.Header().Set("Content-Type", "application/json")

		outprod, err := prod.GetAllProducts(r.Context())
		if err != nil {
			log.Error("failed to get Product", "error", err, "op", op)
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

func GetOneProduct(log *slog.Logger, prod product.OneProductGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.GetOneProduct"
		log.Info("get one poduct", "op", op)

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

		outprod, err := prod.GetOneProduct(r.Context(), id)
		if err != nil {
			log.Error("failed to get Product", "error", err, "op", op)
			http.Error(w, "something went wrong", http.StatusNotFound)
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

func CreateProduct(log *slog.Logger, prod product.ProductCreater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.CreateProduct"
		log.Info("create Product", "op", op)

		if r.Body == nil {
			log.Error("empty body", "op", op)
			http.Error(w, "empty body", http.StatusBadRequest)
			return
		}

		var req domain.Product
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("invalid json", "error", err, "op", op)
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		if req.Name == "" {
			log.Error("empty name", "op", op)
			http.Error(w, "empty name", http.StatusBadRequest)
			return
		}

		err := prod.CreateProduct(r.Context(), req)
		if err != nil {
			log.Error("error updating", "error", err, "op", op)
			http.Error(w, "error updating", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}

}

func UpdateProduct(log *slog.Logger, prod product.ProductUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.UpdateProduct"
		log.Info("update Product", "op", op)

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error("invalid id", "error", err, "op", op)
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		if r.Body == nil {
			log.Error("empty body", "op", op)
			http.Error(w, "empty body", http.StatusBadRequest)
			return
		}
		if id == 0 || id < 0 {
			log.Error("empty id", "op", op)
			http.Error(w, "empty id", http.StatusBadRequest)
			return
		}

		var req domain.Product
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("invalid json", "error", err, "op", op)
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		if req.Name == "" {
			log.Error("empty name", "op", op)
			http.Error(w, "empty name", http.StatusBadRequest)
			return
		}

		req.ID = id

		err = prod.UpdateProduct(r.Context(), req)
		if err != nil {
			log.Error("error updating", "error", err, "op", op)
			http.Error(w, "error updating", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteProduct(log *slog.Logger, prod product.ProductDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.DeleteProduct"
		log.Info("delete Product", "op", op)

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

		err = prod.DelOneProduct(r.Context(), id)
		if err != nil {
			log.Error("error updating", "error", err, "op", op)
			http.Error(w, "error updating", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
