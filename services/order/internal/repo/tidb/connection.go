package tidb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func init() {
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	// 1. Connect to MySQL without specifying the database
	dsnRoot := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort)

	sqlDB, err := sql.Open("mysql", dsnRoot)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL (without db): %v", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Failed to close MySQL connection: %v", err)
		}
	}()

	// 2. Create the database if it does not exist
	_, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS `" + dbName + "`")
	if err != nil {
		log.Fatalf("Failed to create database %s: %v", dbName, err)
	}
	log.Printf("Ensured database `%s` exists", dbName)

	// 3. Seed data for users and merchants if not seeded
	var countUsers int64
	err = sqlDB.QueryRow("SELECT COUNT(*) FROM ztf_db.users").Scan(&countUsers)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Failed to check users table: %v", err)
	}
	if countUsers == 0 {
		seedUsers()
		log.Println("Seeded initial data for users and merchants")
	} else {
		log.Println("Users table already contains data, skipping seeding")
	}

	var countMerchants int64
	err = sqlDB.QueryRow("SELECT COUNT(*) FROM ztf_db.merchants").Scan(&countMerchants)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Failed to check merchants table: %v", err)
	}
	if countMerchants == 0 {
		seedMerchants()
		log.Println("Seeded initial data for merchants")
	} else {
		log.Println("Merchants table already contains data, skipping seeding")
	}
	log.Println("Database initialization complete")
}

func connectDB() *gorm.DB {
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
	return db
}

func GetDB() *gorm.DB {
	once.Do(func() {
		dbInstance = connectDB()
	})

	return dbInstance
}
