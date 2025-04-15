package telegram

import (
	"log"

	"github.com/developeerz/restorio-telegram/config"
	tele "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type bot struct {
	bot   *tele.BotAPI
	cache Cache
}

func NewTelegramBot(repo Repository) *bot {
	botToken := config.ConfigService.BotToken

	bot, err := tele.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	return &bot{
		bot:        bot,
		Repository: repo,
	}
}

func (bot *bot) StartPolling() {
	updateConfig := tele.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "/start":
			bot.cmdStart(&update)

		case "Получить код":
			bot.getCode(&update)
		}
	}
}
