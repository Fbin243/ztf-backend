package composer

import (
	"log"
	"os"
	"ztf-backend/proto/pb/order"
	"ztf-backend/services/promotion/internal/repo/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ComposeOrderClient() (*rpc.OrderClient, *grpc.ClientConn) {
	port, found := os.LookupEnv("ORDER_GPRC_PORT")
	if !found {
		port = "50050"
	}

	host, found := os.LookupEnv("ORDER_GRPC_HOST")
	if !found {
		host = "localhost"
	}

	log.Printf("Connecting to Order gRPC service at %s:%s", host, port)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(host+":"+port, opts)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return rpc.NewOrderClient(order.NewOrderServiceClient(conn)), conn
}
