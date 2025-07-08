package composer

import (
	"log"
	"log/slog"
	"os"
	"ztf-backend/proto/pb/order"
	"ztf-backend/services/promotion/internal/repo/rpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
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

	slog.Info("Connecting to Order gRPC service", "host", host, "port", port)

	conn, err := grpc.NewClient(
		host+":"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return rpc.NewOrderClient(order.NewOrderServiceClient(conn)), conn
}
