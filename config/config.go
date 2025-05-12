package config

import (
	"fmt"
	"os"

	"slices"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	botToken    string
	serviceName string
	brokers     []string
	topic       string
}

func (c *Config) BotToken() string {
	return c.botToken
}

func (c *Config) ServiceName() string {
	return c.serviceName
}

func (c *Config) Brokers() []string {
	return c.brokers
}

func (c *Config) Topic() string {
	return c.topic
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Info().Msgf("load dotenv: %v", err)
	}

	botToken := os.Getenv("BOT_TOKEN")
	serviceName := os.Getenv("SERVICE_NAME")
	broker := os.Getenv("BROKERS")
	topic := os.Getenv("TOPIC")

	if anyIsEmpty(botToken, serviceName, broker, topic) {
		return nil, fmt.Errorf("empty property")
	}

	return &Config{
		botToken:    botToken,
		serviceName: serviceName,
		brokers:     []string{broker},
		topic:       topic,
	}, nil
}

func anyIsEmpty(properties ...string) bool {
	return slices.Contains(properties, "")
}
