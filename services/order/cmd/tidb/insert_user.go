package tidb

import (
	"context"
	"log"
	"os"
	"strconv"
	"ztf-backend/services/order/internal/entity"
	"ztf-backend/services/order/internal/repo/tidb"

	"github.com/brianvoe/gofakeit/v7"
)

func InsertUser(ctx context.Context, userCount int) {
	var entities []entity.User
	for range userCount {
		entities = append(entities, entity.User{
			Username: gofakeit.Username(),
			Email:    gofakeit.Email(),
		})
	}

	ids, err := tidb.NewUserRepo(tidb.GetDB()).InsertMany(
		context.Background(),
		entities,
	)
	if err != nil {
		log.Fatalf("Error inserting users: %v", err)
	}

	// Append the IDs to a file
	filePath := "./tmp/user_ids.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatalf("Error creating user_ids.txt: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}()

	for _, id := range ids {
		_, err := file.WriteString(strconv.FormatInt(id, 10) + "\n")
		if err != nil {
			log.Fatalf("Error writing to user_ids.txt: %v", err)
		}
	}

	log.Printf("Inserted %d users successfully. User IDs appended to %s", userCount, filePath)
}
