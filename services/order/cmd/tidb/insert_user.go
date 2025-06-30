package tidb

import (
	"context"
	"log"
	"os"

	"ztf-backend/pkg/db"
	"ztf-backend/pkg/db/base"
	"ztf-backend/services/order/internal/entity"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/google/uuid"
)

func InsertUser(ctx context.Context, userCount int) {
	var entities []entity.User
	for range userCount {
		entities = append(entities, entity.User{
			BaseEntity: &base.BaseEntity{
				Id: uuid.NewString(),
			},
			Username: gofakeit.Username(),
			Email:    gofakeit.Email(),
		})
	}

	ids, err := base.NewBaseRepo[entity.User](db.GetDB()).InsertMany(
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
	defer file.Close()

	for _, id := range ids {
		_, err := file.WriteString(id + "\n")
		if err != nil {
			log.Fatalf("Error writing to user_ids.txt: %v", err)
		}
	}

	log.Printf("Inserted %d users successfully. User IDs appended to %s", userCount, filePath)
}
