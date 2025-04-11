package main

import (
	"github.com/developeerz/restorio-telegram/config"
	"github.com/developeerz/restorio-telegram/internal/database"
	"github.com/developeerz/restorio-telegram/internal/repository"
	"github.com/developeerz/restorio-telegram/internal/telegram"
)

func main() {
	config.LoadConfig()

	db := database.Connect()
	userRepository := repository.NewRepository(db)

	bot := telegram.NewTelegramBot(userRepository)

	bot.StartPolling()
}
