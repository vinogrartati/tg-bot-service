package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg-bot-service/internal/config"
)

func main() {
	cfg := config.NewConfig()

	bot, err := tgbotapi.NewBotAPI(cfg.TGToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	bot.GetUpdatesChan(u) // TODO переписать на свой метод с несколькими рутинами или на вебхуки
}
