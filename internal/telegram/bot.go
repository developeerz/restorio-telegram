package telegram

import (
	"log"

	"github.com/developeerz/restorio-telegram/config"
	"github.com/developeerz/restorio-telegram/internal/repository"
	tele "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Bot            *tele.BotAPI
	UserRepository repository.UserRepository
}

func NewTelegramBot(r repository.UserRepository) *TelegramBot {
	botToken := config.ConfigService.BotToken
	bot, err := tele.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	return &TelegramBot{
		Bot:            bot,
		UserRepository: r,
	}
}

func (bot *TelegramBot) StartPolling() {
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
