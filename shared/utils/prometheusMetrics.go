package utils

import "github.com/prometheus/client_golang/prometheus"

var (
	EndpointCounterMetrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "api",
			Name:      "http_request_count",
			Help:      "The total number of requests made to some endpoint",
		},
		[]string{"status", "method", "path"},
	)

	EndpointDurationMetrics = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: "api",
			Name:      "http_request_duration_seconds",
			Help:      "Latency of some endpoint requests in seconds",
		},
		[]string{"status", "method", "path"},
	)
)

func init() {
	prometheus.MustRegister(EndpointCounterMetrics)
	prometheus.MustRegister(EndpointDurationMetrics)
}

func PutPrometheusMetrics(path, method, status string) {
	defer func() {
		EndpointCounterMetrics.WithLabelValues(status, method, path).Inc()
	}()
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(_time float64) {
		EndpointDurationMetrics.WithLabelValues(status, method, path).Observe(_time)
	}))
	defer func() {
		timer.ObserveDuration()
	}()
}
