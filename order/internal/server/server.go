package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	biz2 "ztf-backend/order/internal/business"
	repo2 "ztf-backend/order/internal/repo"
	transport2 "ztf-backend/order/internal/transport"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port      int
	orderHdl  *transport2.OrderHandler
	couponHdl *transport2.CouponHandler
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
	couponRepo := repo2.NewCouponRepo()
	orderRepo := repo2.NewOrderRepo()
	userRepo := repo2.NewUserRepo()
	merchantRepo := repo2.NewMerchantRepo()
	orderBusiness := biz2.NewOrderBusiness(orderRepo, couponRepo, userRepo, merchantRepo)
	merchantBusiness := biz2.NewMerchantBusiness(merchantRepo)
	userBusiness := biz2.NewUserBusiness(userRepo)

	NewServer := &Server{
		port: port,
		orderHdl: transport2.NewOrderHandler(
			orderBusiness,
			merchantBusiness,
			userBusiness,
		),
		couponHdl: transport2.NewCouponHandler(
			biz2.NewCouponBusiness(
				couponRepo,
			),
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
