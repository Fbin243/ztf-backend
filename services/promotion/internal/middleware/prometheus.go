package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		HttpRequestSuccess,
		HttpRequestError)
}

// Custom metrics handler with custom registry
func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.HandlerFor(CustomRegistry, promhttp.HandlerOpts{})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Middleware to record incoming requests metrics
func RequestMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// path := c.Request.URL.Path
		c.Next()
		status := c.Writer.Status()
		// observability.HttpRequestDuration.WithLabelValues(strconv.Itoa(status), c.Request.Method, path).Observe(float64(c.Writer.Size()))
		if status < 400 {
			HttpRequestSuccess.WithLabelValues(strconv.Itoa(status)).Inc()
		} else {
			HttpRequestError.WithLabelValues(strconv.Itoa(status)).Inc()
		}
	}
}
