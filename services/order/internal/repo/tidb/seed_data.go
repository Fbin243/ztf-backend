package tidb

import (
	"context"
	"log"
	"ztf-backend/services/order/internal/entity"
)

func seedUsers() {
	ids, err := NewUserRepo(GetDB()).InsertMany(context.Background(), []entity.User{
		{
			Id:       1,
			Username: "ntbinh",
			Email:    "ntbinh243@gmail.com",
		},
		{
			Id:       2,
			Username: "dqtrieu",
			Email:    "dqtrieu@gmail.com",
		},
		{
			Id:       3,
			Username: "hatra",
			Email:    "hatra@gmail.com",
		},
	})
	if err != nil {
		log.Printf("Failed to seed users: %v", err)
	}

	log.Printf("Inserted users with IDs: %v", ids)
}

func seedMerchants() {
	merchantIds, err := NewMerchantRepo(GetDB()).InsertMany(context.Background(), []entity.Merchant{
		{
			Id:       4,
			Username: "highland",
			Email:    "merchant@highland.com",
		},
		{
			Id:       5,
			Username: "phuclong",
			Email:    "merchant2@phuclong.com",
		},
		{
			Id:       6,
			Username: "thecoffeehouse",
			Email:    "merchant3@thecoffeehouse.com",
		},
	})
	if err != nil {
		log.Printf("Failed to seed merchants: %v", err)
	}

	log.Printf("Inserted merchants with IDs: %v", merchantIds)
}
