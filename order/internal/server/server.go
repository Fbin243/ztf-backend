package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"ztf-backend/order/internal/composer"
	"ztf-backend/order/internal/transport"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
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

	port, _ := strconv.Atoi(os.Getenv("ORDER_PORT"))
	log.Printf("Starting server on port %d", port)

	composer := composer.GetComposer()
	NewServer := &Server{
		port: port,
		orderHdl: transport.NewOrderHandler(
			composer.OrderBusiness,
			composer.MerchantBusiness,
			composer.UserBusiness,
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
