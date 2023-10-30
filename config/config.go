package config

import (
	"api-donasi/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
		log.Fatalf("Cannot load .env file. Err: %s", err)
	}

	InitDB()
	InitialMigration()  // User
	InitialMigration2() // Campaign
	InitialMigration3() // Donation
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	config := Config{
		// DB_Username: "DB_USER",
		// DB_Password: "DB_PASSWORD",
		// DB_Port:     "DB_PORT",
		// DB_Host:     "DB_HOST",
		// DB_Name:     "DB_NAME",
		DB_Username: os.Getenv("DB_USER"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Name:     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
}

func InitialMigration2() {
	DB.AutoMigrate(&models.Campaign{})
}

func InitialMigration3() {
	DB.AutoMigrate(&models.Donation{})
}
