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
	"ztf-backend/services/order/internal/composer"
	"ztf-backend/services/order/internal/transport/rest"
)

type Server struct {
	port     int
	orderHdl *rest.OrderHandler
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
		orderHdl: rest.NewOrderHandler(
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
