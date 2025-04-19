package telegram

import (
	"context"
	"fmt"

	"github.com/developeerz/restorio-telegram/config"
	tele "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type Bot struct {
	bot   *tele.BotAPI
	cache Cache
}

func NewTelegramBot(cache Cache) (*Bot, error) {
	botToken := config.ConfigService.BotToken

	bot, err := tele.NewBotAPI(botToken)
	if err != nil {
		return nil, fmt.Errorf("new bot: %w", err)
	}

	return &Bot{bot: bot, cache: cache}, nil
}

func (bot *Bot) StartPolling(ctx context.Context) {
	var err error

	updateConfig := tele.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "/start":
			err = bot.cmdStart(&update)
			if err != nil {
				log.Error().AnErr("error", err).Send()
			}

		case "Получить код":
			err = bot.getCode(ctx, &update)
			if err != nil {
				log.Error().AnErr("error", err).Send()
			}
		}
	}
}
