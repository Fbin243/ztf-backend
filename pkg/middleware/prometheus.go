package middleware

import (
	"strconv"

	"ztf-backend/pkg/observability"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Custom metrics handler with custom registry
func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.HandlerFor(observability.CustomRegistry, promhttp.HandlerOpts{})
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
			observability.HttpRequestSuccess.WithLabelValues(strconv.Itoa(status)).Inc()
		} else {
			observability.HttpRequestError.WithLabelValues(strconv.Itoa(status)).Inc()
		}
	}
}
