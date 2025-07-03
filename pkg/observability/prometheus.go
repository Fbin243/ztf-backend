package observability

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Define metrics
var (
	HttpRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_request_total",
		Help: "Total number of requests",
	}, []string{"status"})

	HttpRequestSuccess = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_request_success",
		Help: "Number of successful requests",
	}, []string{"status"})

	HttpRequestError = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_request_error",
		Help: "Number of failed requests",
	}, []string{"status"})

	HttpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "api_http_request_duration_seconds",
		Help:    "Duration of HTTP requests in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"status", "method", "path"})
)

// Custom registry (without default Go metrics)
var CustomRegistry = prometheus.NewRegistry()

// Register metrics with custom registry
func init() {
	CustomRegistry.MustRegister(
		prometheus.NewGoCollector(),
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		HttpRequestSuccess,
		HttpRequestError)
}
