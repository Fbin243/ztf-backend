package middleware

import (
	"net/http"
	"ztf-backend/services/order/internal/auth"

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

		ctx.Set(string(auth.AuthKey), userID)
		ctx.Next()
	}
}
