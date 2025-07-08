package middleware

import (
	"log/slog"
	"net/http"
	"ztf-backend/services/order/internal/errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Step 1: Process the request first

		// Step 2: Check if any errors were added to the context
		if len(c.Errors) > 0 {
			var customErr *errors.Error
			for _, e := range c.Errors {
				if err, ok := e.Err.(*errors.Error); ok {
					customErr = err
				} else {
					customErr = &errors.Error{
						Message: e.Err.Error(),
						Code:    errors.ErrCodeInternalError,
					}
				}

				// Step 3: Log the error
				slog.Error("Rest API error",
					slog.String("method", c.Request.Method),
					slog.String("path", c.Request.URL.Path),
					slog.String("client_ip", c.ClientIP()),
					slog.String("error", customErr.Message),
					slog.String("code", string(customErr.Code)),
				)
			}

			// Step 4. Respond with the last error
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    customErr.Code,
				"message": customErr.Message,
			})
		}
	}
}
