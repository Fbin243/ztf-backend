package tidb

import (
	"context"
	"log"
	"ztf-backend/services/order/internal/entity"
)

func seedUsers() {
	ids, err := NewUserRepo(GetDB()).InsertMany(context.Background(), []entity.User{
		{
			Id:       "a1927cb1-1db0-4b18-91ed-578559ba7489",
			Username: "ntbinh",
			Email:    "ntbinh243@gmail.com",
		},
		{
			Id:       "bb8a677a-c9f3-46ca-8299-54b23d2c4d23",
			Username: "dqtrieu",
			Email:    "dqtrieu@gmail.com",
		},
		{
			Id:       "a7dacaa5-baa2-4458-b62e-27a5a073dfb1",
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
			Id:       "53c7e139-5c92-49e7-a4b2-667782e8fd9e",
			Username: "highland",
			Email:    "merchant@highland.com",
		},
		{
			Id:       "1540bf4a-07d6-48b9-8047-726c9150cf1f",
			Username: "phuclong",
			Email:    "merchant2@phuclong.com",
		},
		{
			Id:       "9b66ffae-02d1-48b0-94e3-937adf52f85a",
			Username: "thecoffeehouse",
			Email:    "merchant3@thecoffeehouse.com",
		},
	})
	if err != nil {
		log.Printf("Failed to seed merchants: %v", err)
	}

	log.Printf("Inserted merchants with IDs: %v", merchantIds)
}
