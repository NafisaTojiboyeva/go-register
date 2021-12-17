package config

import (
	"os"
	"log"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
)

func DB() (*gorm.DB, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PG_HOST := os.Getenv("PG_HOST")
	PG_USER := os.Getenv("PG_USER")
	PG_PASSWORD := os.Getenv("PG_PASSWORD")
	PG_PORT := os.Getenv("PG_PORT")
	PG_DBNAME := os.Getenv("PG_DBNAME")
	PG_SSLMODE := os.Getenv("PG_SSLMODE")
	PG_TIMEZONE := os.Getenv("PG_TIMEZONE")

	PGConnection := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=%s",
		PG_HOST,
		PG_USER,
		PG_PASSWORD,
		PG_PORT,
		PG_DBNAME,
		PG_SSLMODE,
		PG_TIMEZONE,
	)

	return gorm.Open(postgres.Open(PGConnection), &gorm.Config{})
}