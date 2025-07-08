package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func ErrorLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				slog.Error("API Error",
					slog.String("method", c.Request.Method),
					slog.String("path", c.Request.URL.Path),
					slog.String("client_ip", c.ClientIP()),
					slog.String("error", e.Error()),
					slog.Int("status", c.Writer.Status()),
				)
			}
		}
	}
}
