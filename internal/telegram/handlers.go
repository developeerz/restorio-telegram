package telegram

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/developeerz/restorio-auth/pkg/repository/redis"
	tele "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *Bot) cmdStart(update *tele.Update) error {
	msg := tele.NewMessage(update.Message.Chat.ID, "Добро пожаловать в Restorio!")
	msg.ReplyMarkup = getMainMenuKeyboard()

	_, err := bot.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (bot *Bot) getCode(ctx context.Context, update *tele.Update) error {
	var user redis.User
	telegram := update.Message.From.UserName
	telegramID := update.Message.From.ID

	code, err := bot.cache.GetVerificationCode(ctx, telegram)
	if err != nil {
		msg := tele.NewMessage(update.Message.Chat.ID, "Вы уже зарегистрированы!")

		_, err = bot.bot.Send(msg)
		if err != nil {
			return fmt.Errorf("bot send: %w", err)
		}

		msg.ReplyMarkup = getMainMenuKeyboard()

		return fmt.Errorf("cannot get verivication code with telegram (%s)", telegram)
	}

	userBytes, err := bot.cache.GetUser(ctx, telegram)
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}

	user.TelegramID = telegramID

	userBytesUpd, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	err = bot.cache.PutUser(ctx, telegram, userBytesUpd)
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	msg := tele.NewMessage(update.Message.Chat.ID, "Ваш код:")

	_, err = bot.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("bot send: %w", err)
	}

	msg = tele.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%d", code))
	msg.ReplyMarkup = getMainMenuKeyboard()

	_, err = bot.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("bot send: %w", err)
	}

	return nil
}
