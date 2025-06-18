package mwPrometheus

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	metric "github.com/go-mod-test/goods/internal/metrics"
)

func Prometheus(log *slog.Logger, m *metric.Prometheus) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			duration := time.Since(start).Seconds()

			route := r.URL.Path
			if ctx := chi.RouteContext(r.Context()); ctx != nil && ctx.RoutePattern() != "" {
				route = ctx.RoutePattern()
			}

			m.HTTPRequestsTotal.WithLabelValues(
				r.Method, route, strconv.Itoa(ww.Status()),
			).Inc()

			m.HTTPRequestsDuration.WithLabelValues(
				r.Method, route,
			).Observe(duration)
		})
	}
}
