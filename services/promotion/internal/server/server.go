package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"ztf-backend/services/order/pkg/observability"
	"ztf-backend/services/promotion/internal/composer"
	"ztf-backend/services/promotion/internal/transport/rest"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port         int
	promotionHdl *rest.PromotionHandler
}

func NewServer() *http.Server {
	// Load environment variables
	appEnv := "dev"
	err := godotenv.Load(".env." + appEnv)
	if err != nil {
		fmt.Printf("Error loading .env.%s file: %v\n", appEnv, err)
	}

	port, _ := strconv.Atoi(os.Getenv("PROMOTION_PORT"))
	log.Printf("Starting server on port %d", port)

	err = observability.Init("promotion-service")
	if err != nil {
		log.Fatalf("Failed to init observability: %v", err)
	}

	composer := composer.GetComposer()
	NewServer := &Server{
		port:         port,
		promotionHdl: rest.NewPromotionHandler(composer.PromotionBusiness),
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
