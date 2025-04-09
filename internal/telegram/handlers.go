package telegram

import (
	"fmt"
	"log"

	tele "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *Bot) cmdStart(update *tele.Update) {
	msg := tele.NewMessage(update.Message.Chat.ID, "Добро пожаловать в Restorio!")
	msg.ReplyMarkup = getMainMenuKeyboard()

	_, err := bot.Bot.Send(msg)
	if err != nil {
		log.Printf("cmdStart(): %v", err)
	}
}

func (bot *Bot) getCode(update *tele.Update) {
	telegram := update.Message.From.UserName
	telegramID := update.Message.From.ID

	code, err := bot.Repository.GetCodeByTelegram(telegram)
	if err != nil {
		msg := tele.NewMessage(update.Message.Chat.ID, "Вы уже зарегистрированы!")

		_, err = bot.Bot.Send(msg)
		if err != nil {
			log.Printf("getCode(%d): %v", telegramID, err)
		}

		msg.ReplyMarkup = getMainMenuKeyboard()
	}

	err = bot.Repository.UpdateUserByTelegram(telegram, telegramID)
	if err != nil {
		log.Printf("UpdateUserByTelegram(%s, %d): %v", telegram, telegramID, err)
	}

	msg := tele.NewMessage(update.Message.Chat.ID, "Ваш код:")

	_, err = bot.Bot.Send(msg)
	if err != nil {
		log.Printf("getCode(%d): %v", telegramID, err)
	}

	msg = tele.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%d", code))
	msg.ReplyMarkup = getMainMenuKeyboard()

	_, err = bot.Bot.Send(msg)
	if err != nil {
		log.Printf("getCode(%d): %v", telegramID, err)
	}
}
