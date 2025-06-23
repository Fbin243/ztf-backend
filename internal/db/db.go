package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"ztf-backend/internal/entity"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func createDB() *gorm.DB {
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Printf("Connected to database %s at %s:%s", dbName, dbHost, dbPort)

	// Auto-migrate for development purposes
	if err := db.AutoMigrate(&entity.Order{}); err != nil {
		log.Fatalf("Failed to auto-migrate Order entity: %v", err)
	}

	return db
}

func GetDB() *gorm.DB {
	once.Do(func() {
		dbInstance = createDB()
	})

	return dbInstance
}
