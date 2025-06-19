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

	// coupon routes
	r.GET("/api/v1/coupons", gin.HandlerFunc(s.couponHdl.GetAllCoupons))
	r.GET("/api/v1/coupons/:id", gin.HandlerFunc(s.couponHdl.GetCouponById))
	r.POST("/api/v1/coupons", gin.HandlerFunc(s.couponHdl.CreateCoupon))
	r.PUT("/api/v1/coupons/:id", gin.HandlerFunc(s.couponHdl.UpdateCoupon))
	r.DELETE("/api/v1/coupons/:id", gin.HandlerFunc(s.couponHdl.DeleteCoupen))

	// order routes
	r.GET("/api/v1/orders", gin.HandlerFunc(s.orderHdl.GetAllOrders))
	r.GET("/api/v1/orders/:id", gin.HandlerFunc(s.orderHdl.GetOrderById))
	r.POST("/api/v1/orders", gin.HandlerFunc(s.orderHdl.CreateOrder))
	r.PUT("/api/v1/orders/:id", gin.HandlerFunc(s.orderHdl.UpdateOrder))
	r.DELETE("/api/v1/orders/:id", gin.HandlerFunc(s.orderHdl.DeleteOrder))

	return r
}
