package telegram

import tele "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func getMainMenuKeyboard() tele.ReplyKeyboardMarkup {
	return tele.NewReplyKeyboard(
		tele.NewKeyboardButtonRow(
			tele.NewKeyboardButton("Получить код"),
		),
	)
}
