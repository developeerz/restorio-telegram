package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	BotToken    string
	ServiceName string
}

var ConfigService Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Info().Msgf("load dotenv: %v", err)
	}

	ConfigService = Config{BotToken: os.Getenv("BOT_TOKEN")}
}
