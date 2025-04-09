package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Postgres string
	BotToken string
}

var ConfigService Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("load dotenv: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	ConfigService = Config{
		Postgres: dsn,
		BotToken: os.Getenv("BOT_TOKEN"),
	}
}
