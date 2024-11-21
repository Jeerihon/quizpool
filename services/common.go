package services

import (
	"fmt"
	"github.com/Jeerihon/quizpool/constants"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я бот помощник по Квиз Плиз.\nЯ могу помочь тебе узнать о предстоящих играх.")
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func Help(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var commandList string
	for _, command := range constants.BotCommands {
		commandList += fmt.Sprintf("/%s - %s\n", command.Command, command.Description)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Вот список доступных команд:\n\n%s", commandList))
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
