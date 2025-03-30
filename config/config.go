package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Postgres string
	BotToken string
}

var ConfigService Config

func LoadConfig() {
	godotenv.Load()

	ConfigService = Config{
		Postgres: os.Getenv("POSTGRES_CONFIG"),
		BotToken: os.Getenv("BOT_TOKEN"),
	}
}
