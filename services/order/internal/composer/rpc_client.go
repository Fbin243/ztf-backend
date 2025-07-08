package composer

import (
	"log"
	"log/slog"
	"os"
	"ztf-backend/proto/pb/promotion"
	"ztf-backend/services/order/internal/repo/rpc"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	slog.Info("Connecting to Promotion gRPC service", "host", host, "port", port)

	conn, err := grpc.NewClient(
		host+":"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return rpc.NewPromotionClient(promotion.NewPromotionServiceClient(conn)), conn
}
