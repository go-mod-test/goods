package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Prometheus struct {
	HTTPRequestsTotal    *prometheus.CounterVec
	HTTPRequestsDuration *prometheus.HistogramVec
}

func NewPrometheus() *Prometheus {
	return &Prometheus{
		HTTPRequestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests.",
			},
			[]string{"method", "path", "status"},
		),
		HTTPRequestsDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "http_requests_duration_seconds",
				Help: "Duration of HTTP requests.",
			},
			[]string{"method", "path"},
		),
	}
}

func (m *Prometheus) RegisterAll() {
	prometheus.MustRegister(m.HTTPRequestsTotal, m.HTTPRequestsDuration)
}
