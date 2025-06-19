package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	biz "ztf-backend/internal/business"
	"ztf-backend/internal/repo"
	"ztf-backend/internal/transport"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port          int
	orderhandler  *transport.OrderHandler
	couponhandler *transport.CouponHandler
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
	NewServer := &Server{
		port:          port,
		orderhandler:  transport.NewOrderHandler(biz.NewOrderBusiness(repo.NewOrderRepo())),
		couponhandler: transport.NewCouponHandler(biz.NewCouponBusiness(repo.NewCouponRepo())),
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
