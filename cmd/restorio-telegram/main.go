package main

import (
	"context"

	"github.com/developeerz/restorio-telegram/config"
	"github.com/developeerz/restorio-telegram/internal/kafka"
	"github.com/developeerz/restorio-telegram/internal/telegram"
	"github.com/rs/zerolog/log"

	logger "github.com/developeerz/restorio-auth/pkg/logger"
	redis "github.com/developeerz/restorio-auth/pkg/repository/redis"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().AnErr("error", err).Send()
	}

	err = logger.InitLogger(cfg.ServiceName())
	if err != nil {
		log.Fatal().AnErr("error", err).Send()
	}

	rdb, err := redis.Connect()
	if err != nil {
		log.Fatal().AnErr("error", err).Send()
	}

	userCache := redis.NewUserCache(rdb)

	bot, err := telegram.NewTelegramBot(cfg.BotToken(), userCache)
	if err != nil {
		log.Fatal().AnErr("error", err).Send()
	}

	kafka := kafka.New(bot, cfg.Brokers(), cfg.Topic())

	go kafka.ReadLoop(ctx)
	bot.StartPolling(ctx)
}
