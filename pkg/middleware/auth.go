package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetHeader("X-User-Id")
		if userID == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("userID", userID)
		ctx.Next()
	}
}
