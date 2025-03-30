package telegram

import (
	"fmt"
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
			msg := tele.NewMessage(update.Message.Chat.ID, "Добро пожаловать в Restorio!")
			bot.Bot.Send(msg)

		case "Получить код":
			telegram := update.Message.From.UserName
			telegramID := update.Message.From.ID

			code, err := bot.UserRepository.GetCodeByTelegram(telegram)
			if err != nil {
				msg := tele.NewMessage(update.Message.Chat.ID, "Вы уже зарегистрированы!")
				bot.Bot.Send(msg)
				continue
			}

			bot.UserRepository.UpdateUserByTelegram(telegram, telegramID)
			msg := tele.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Ваш код: %d", code))
			bot.Bot.Send(msg)
		}
	}
}
