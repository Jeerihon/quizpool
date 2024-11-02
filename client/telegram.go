package client

import (
	"github.com/Jeerihon/quizpool/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Init() *tgbotapi.BotAPI {
	cfg := config.LoadConfig()

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.APIToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}
