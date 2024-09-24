package database

import (
	"employeeregister/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	timeZone := os.Getenv("DB_TIMEZONE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host, user, password, dbname, port, timeZone)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database successfully")
	DB.AutoMigrate(&models.Employee{})
	DB.AutoMigrate(&models.Warranty{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.WrrantyData{})
	DB.AutoMigrate(&models.Otp{})
	DB.AutoMigrate(&models.Votes{})
	DB.AutoMigrate(&models.Candidate{})
	DB.AutoMigrate(&models.UserMessage{})
}
