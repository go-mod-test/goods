package main

/*
func main() {
	config := config.MustLoadEnvConfig()

	logger := mylogger.SetupLogger(config.Env)

	ctxApp, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	router := chi.NewRouter()

	m := metrics.NewPrometheus()

	logger.Info(
		"Starting server",
		slog.String("env", config.Env),
	)
	logger.Debug("Debug message is enabled")

	db, err := storage.NewStorage(config.Db)
	if err != nil {
		logger.Error("Error connecting to database:%s", "error", err)
		os.Exit(1)
	}
	defer db.Db.Close()

	wordDb := postgres.NewWordStorage(logger, db.Db)
	logger.Debug("Connected to database")

	//
	// mideleware
	router.Use(mwPrometheus.Prometheus(logger, m))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(mwLogger.New(logger))

	//alias
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	router.Get("/metrics", promhttp.Handler().ServeHTTP)
	router.Get("/api/get/{word}", handlers.GetWord(logger, wordDb))
	router.Post("/api/add", handlers.CreateWord(logger, wordDb))
	router.Delete("/api/delete/{word}", handlers.DeleteWord(logger, wordDb))
	router.Put("/api/update/{word}", handlers.UpdateWord(logger, wordDb))
	//

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
*/
