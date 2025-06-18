package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-mod-test/goods/internal/config"
	"github.com/go-mod-test/goods/internal/handlers"
	mylogger "github.com/go-mod-test/goods/internal/logger"
	metric "github.com/go-mod-test/goods/internal/metrics"
	mwPrometheus "github.com/go-mod-test/goods/internal/middleware/prometheus"
	"github.com/go-mod-test/goods/internal/storage"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	//config
	config := config.MustLoadEnvConfig()

	//logger
	logger := mylogger.SetupLogger(config.Env)

	logger.Info(
		"Starting server",
		slog.String("env", config.Env),
	)
	logger.Debug("Debug message is enabled")

	//ctx
	ctxApp, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// db
	db, err := storage.NewStorage(config.Db)
	if err != nil {
		logger.Error("Error connecting to database:%s", "error", err)
		os.Exit(1)
	}
	defer db.Db.Close()

	//promitheus

	m := metric.NewPrometheus()

	//router
	router := chi.NewRouter()

	router.Use(mwPrometheus.Prometheus(logger, m))

	//aliases
	router.Get("/metrics", promhttp.Handler().ServeHTTP)
	//customer
	router.Get("/customer/all", handlers.GetAllCustomers(logger, db))
	router.Get("/customer/{id}", handlers.GetOneCustomer(logger, db))
	router.Post("/customer/create", handlers.CreateCustomer(logger, db))
	router.Put("/customer/update/{id}", handlers.UpdateCustomer(logger, db))
	router.Delete("/customer/delete/{id}", handlers.DeleteCustomer(logger, db))

	//product
	router.Get("/product/all", handlers.GetAllProducts(logger, db))
	router.Get("/product/{id}", handlers.GetOneProduct(logger, db))
	router.Post("/product/create", handlers.CreateProduct(logger, db))
	router.Put("/product/update/{id}", handlers.UpdateProduct(logger, db))
	router.Delete("/product/delete/{id}", handlers.DeleteProduct(logger, db))

	//invoice
	router.Get("/invoice/all", handlers.GetAllInvoices(logger, db))
	router.Get("/invoice/{id}", handlers.GetOneInvoice(logger, db))
	router.Post("/invoice/create", handlers.CreateInvoice(logger, db))
	router.Put("/invoice/update/{id}", handlers.UpdateInvoice(logger, db))
	router.Delete("/invoice/delete/{id}", handlers.DeleteInvoice(logger, db))

	//invoiceItem
	router.Get("/invoiceItem/{id}", handlers.GetOneInvoiceItem(logger, db))
	router.Post("/invoiceItem/create", handlers.CreateInvoiceItem(logger, db))
	router.Delete("/invoiceItem/delete/{id}", handlers.DeleteInvoiceItem(logger, db))

	//server
	srv := &http.Server{
		Addr:         config.HTTPServer.Host,
		Handler:      router,
		ReadTimeout:  config.HTTPServer.ReadTimeout,
		WriteTimeout: config.HTTPServer.WriteTimeout,
		IdleTimeout:  config.HTTPServer.IdleTimeout,
	}

	serverDone := make(chan struct{})

	go func() {
		defer close(serverDone)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("failed to start server", "error", err)
		}
	}()

	logger.Info("SERVER START!!!!")

	go func() {
		<-ctxApp.Done()
		logger.Info("shutting down...")
		ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctxShutdown); err != nil {
			logger.Error("failed to shutdown server", "error", err)
		}
	}()

	<-serverDone
	logger.Info("SERVER STOPED")
}
