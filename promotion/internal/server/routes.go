package server

import (
	"net/http"

	"ztf-backend/pkg/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-User-Id"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.Use(middleware.AuthMiddleware())

	// promotion routes
	r.GET("/api/v1/promotions/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
	r.GET("/api/v1/promotions", s.promotionHdl.GetAllPromotions)
	r.GET("/api/v1/promotions/search", s.promotionHdl.GetPromotionByCode)
	r.GET("/api/v1/promotions/:id", s.promotionHdl.GetPromotionById)
	r.POST("/api/v1/promotions", s.promotionHdl.CreatePromotion)
	r.PUT("/api/v1/promotions/:id", s.promotionHdl.UpdatePromotion)
	r.DELETE("/api/v1/promotions/:id", s.promotionHdl.DeletePromotion)
	r.POST("/api/v1/promotions/verify", s.promotionHdl.VerifyPromotion)
	r.POST("/api/v1/promotions/collect/:id", s.promotionHdl.CollectPromotion)

	return r
}
