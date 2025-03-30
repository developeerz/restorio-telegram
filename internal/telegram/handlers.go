package telegram

import (
	"fmt"

	tele "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *TelegramBot) cmdStart(update *tele.Update) {
	msg := tele.NewMessage(update.Message.Chat.ID, "Добро пожаловать в Restorio!")
	msg.ReplyMarkup = getMainMenuKeyboard()
	bot.Bot.Send(msg)
}

func (bot *TelegramBot) getCode(update *tele.Update) {
	telegram := update.Message.From.UserName
	telegramID := update.Message.From.ID

	code, err := bot.UserRepository.GetCodeByTelegram(telegram)
	if err != nil {
		msg := tele.NewMessage(update.Message.Chat.ID, "Вы уже зарегистрированы!")
		bot.Bot.Send(msg)
		msg.ReplyMarkup = getMainMenuKeyboard()
		return
	}

	bot.UserRepository.UpdateUserByTelegram(telegram, telegramID)

	msg := tele.NewMessage(update.Message.Chat.ID, "Ваш код:")
	bot.Bot.Send(msg)
	msg = tele.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%d", code))
	msg.ReplyMarkup = getMainMenuKeyboard()
	bot.Bot.Send(msg)
}
