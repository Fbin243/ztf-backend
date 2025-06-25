package composer

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"ztf-backend/order/internal/repo/rpc"
	"ztf-backend/proto/pb/promotion"
)

func ComposePromotionClient() (*rpc.PromotionClient, *grpc.ClientConn) {
	port, found := os.LookupEnv("PROMOTION_GPRC_PORT")
	if !found {
		port = "50051"
	}

	host, found := os.LookupEnv("PROMOTION_GRPC_HOST")
	if !found {
		host = "localhost"
	}

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(host+":"+port, opts)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return rpc.NewPromotionClient(promotion.NewPromotionServiceClient(conn)), conn
}
