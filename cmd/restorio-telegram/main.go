package main

import (
	"github.com/developeerz/restorio-telegram/config"
	"github.com/developeerz/restorio-telegram/internal/telegram"
	"github.com/rs/zerolog/log"

	logger "github.com/developeerz/restorio-auth/pkg/logger"
	redis "github.com/developeerz/restorio-auth/pkg/repository/redis"
)

func main() {
	config.LoadConfig()

	err := logger.InitLogger(config.ConfigService.ServiceName)
	if err != nil {
		log.Fatal().AnErr("error", err).Send()
	}

	rdb, err := redis.Connect()
	if err != nil {
		log.Fatal().AnErr("error", err).Send()
	}

	userCache := redis.NewUserCache(rdb)

	bot, err := telegram.NewTelegramBot(userCache)
	if err != nil {
		log.Fatal().AnErr("error", err).Send()
	}

	bot.StartPolling()
}
