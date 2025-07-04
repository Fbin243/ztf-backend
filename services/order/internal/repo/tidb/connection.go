package tidb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func init() {
	// Load environment variables
	appEnv := "dev"
	err := godotenv.Load(".env." + appEnv)
	if err != nil {
		fmt.Printf("Error loading .env.%s file: %v\n", appEnv, err)
	}

	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	// 1. Connect to MySQL without specifying the database
	dsnRoot := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort)
	log.Printf("dsnRoot: %s", dsnRoot)

	sqlDB, err := sql.Open("mysql", dsnRoot)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
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

	// 3. Create tables and seed data for users and merchants
	_, err = sqlDB.Exec(`
		CREATE TABLE
		IF NOT EXISTS ztf_db.users (
			id CHAR(36) PRIMARY KEY,
			username VARCHAR(255) NOT NULL UNIQUE,
			email VARCHAR(255) NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
	seedUsers()

	_, err = sqlDB.Exec(`
		CREATE TABLE
		IF NOT EXISTS ztf_db.merchants (
			id CHAR(36) PRIMARY KEY,
			username VARCHAR(255) NOT NULL UNIQUE,
			email VARCHAR(255) NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatalf("Failed to create merchants table: %v", err)
	}
	seedMerchants()
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
