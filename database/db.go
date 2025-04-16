package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDatabase() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	envdata := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("DB_NAME"), os.Getenv("PASSWORD"))

	var err error
	Db, err = sql.Open("postgres", envdata)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	fmt.Println("Successfully connected to database!")
}
