package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connect *gorm.DB

func init() {
	fmt.Println("Connecting to database...")

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		host,
		user,
		password,
		dbname,
		port,
	)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	connect = db.Debug()

	fmt.Println("Connected to database successfully!")
}

func GetConnection() *gorm.DB { return connect }
