package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	// order routes
	r.GET("/api/v1/orders/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
	r.GET("/api/v1/orders", s.orderHdl.GetAllOrders)
	r.GET("/api/v1/orders/:id", s.orderHdl.GetOrderById)
	r.POST("/api/v1/orders", s.orderHdl.CreateOrder)
	r.PUT("/api/v1/orders/:id", s.orderHdl.UpdateOrder)
	r.PUT("/api/v1/orders/:id/pay", s.orderHdl.PayForOrder)
	r.DELETE("/api/v1/orders/:id", s.orderHdl.DeleteOrder)

	return r
}
