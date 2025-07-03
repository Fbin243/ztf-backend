package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"ztf-backend/services/order/cmd/tidb"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			&tidb.InsertUserCmd,
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			// Load environment variables
			appEnv := "dev"
			err := godotenv.Load(".env." + appEnv)
			if err != nil {
				fmt.Printf("Error loading .env.%s file: %v\n", appEnv, err)
			}

			return ctx, nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
