package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresqlConnection() *gorm.DB {
	fmt.Println("Connecting to database...")

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

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

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	fmt.Println("Connected to database successfully!")

	db = db.Debug()

	return db
}
