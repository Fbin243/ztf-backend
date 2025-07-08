package server

import (
	"net/http"
	"ztf-backend/services/order/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-User-Id"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	// Register /metrics before middleware to avoid couting the request from prometheus to /metrics
	r.GET("/metrics", middleware.PrometheusHandler())
	r.Use(middleware.RequestMetricsMiddleware())
	r.Use(otelgin.Middleware("order-service"))
	r.Use(middleware.ErrorHandler())

	r.Use(middleware.AuthHandler())

	// order routes
	r.GET("/api/v1/orders", s.orderHdl.GetAllOrders)
	r.GET("/api/v1/orders/:id", s.orderHdl.GetOrderById)
	r.POST("/api/v1/orders", s.orderHdl.CreateOrder)
	r.PUT("/api/v1/orders/:id", s.orderHdl.UpdateOrder)
	r.PUT("/api/v1/orders/pay/:id", s.orderHdl.PayForOrder)
	r.DELETE("/api/v1/orders/:id", s.orderHdl.DeleteOrder)

	return r
}
