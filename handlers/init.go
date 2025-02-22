package handlers

import (
	"github.com/Jeerihon/quizpool/constants"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			Callbacks(bot, update)
		} else if update.Message != nil && update.Message.IsCommand() {
			Commands(bot, update)
		}
	}
}

func InitCommands(bot *tgbotapi.BotAPI) {
	_, err := bot.Request(tgbotapi.SetMyCommandsConfig{
		Commands: constants.BotCommands,
	})
	if err != nil {
		panic(err)
	}
}
