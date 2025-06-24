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
	biz "ztf-backend/order/internal/business"
	"ztf-backend/order/internal/repo"
	"ztf-backend/order/internal/transport"
)

type Server struct {
	port     int
	orderHdl *transport.OrderHandler
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
	orderRepo := repo.NewOrderRepo()
	userRepo := repo.NewUserRepo()
	merchantRepo := repo.NewMerchantRepo()
	orderBusiness := biz.NewOrderBusiness(orderRepo, userRepo, merchantRepo)
	merchantBusiness := biz.NewMerchantBusiness(merchantRepo)
	userBusiness := biz.NewUserBusiness(userRepo)

	NewServer := &Server{
		port: port,
		orderHdl: transport.NewOrderHandler(
			orderBusiness,
			merchantBusiness,
			userBusiness,
		),
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
