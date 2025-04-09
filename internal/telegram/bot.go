package telegram

import (
	"log"

	"github.com/developeerz/restorio-telegram/config"
	tele "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	Bot        *tele.BotAPI
	Repository Repository
}

func NewTelegramBot(repo Repository) *Bot {
	botToken := config.ConfigService.BotToken

	bot, err := tele.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	return &Bot{
		Bot:        bot,
		Repository: repo,
	}
}

func (bot *Bot) StartPolling() {
	updateConfig := tele.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.Bot.GetUpdatesChan(updateConfig)

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
