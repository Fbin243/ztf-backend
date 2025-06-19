package server

import (
	"net/http"
	"time"

	"ztf-backend/internal/db"
	"ztf-backend/internal/entity"

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

	r.GET("/", s.HelloWorldHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	tidb := db.GetDB()
	coupon := entity.Coupon{
		Code:           "TESTCOUPON",
		Name:           "Test Coupon",
		Description:    "This is a test coupon",
		CouponType:     entity.CouponTypePercentage,
		UsageMethod:    entity.UsageMethodSingleUse,
		ExpirationDate: time.Now().AddDate(0, 1, 0),
	}
	tidb.Save(&coupon)

	c.JSON(http.StatusOK, coupon)
}
