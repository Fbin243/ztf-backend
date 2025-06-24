package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	biz "ztf-backend/promotion/internal/business"
	"ztf-backend/promotion/internal/repo"
	"ztf-backend/promotion/internal/transport"
)

type Server struct {
	port         int
	promotionHdl *transport.PromotionHandler
}

func NewServer() *http.Server {
	// Load environment variables
	appEnv := "dev"
	err := godotenv.Load(".env." + appEnv)
	if err != nil {
		fmt.Printf("Error loading .env.%s file: %v\n", appEnv, err)
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	log.Printf("Starting server on port %d", port)

	// Dependency injection
	promotionRepo := repo.NewPromotionRepo()
	promotionBusiness := biz.NewPromotionBusiness(promotionRepo)
	promotionHandler := transport.NewPromotionHandler(promotionBusiness)

	NewServer := &Server{
		port:         port,
		promotionHdl: promotionHandler,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
