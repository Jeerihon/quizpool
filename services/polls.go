package services

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateAttendancePoll(bot *tgbotapi.BotAPI, update tgbotapi.Update, quizEventId string) {
	poll := tgbotapi.SendPollConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           update.CallbackQuery.Message.Chat.ID,
			ReplyToMessageID: update.CallbackQuery.Message.MessageID,
		},
		Options: []string{
			"Иду",
			"Не иду",
			"Скорее иду",
			"Скорее не иду",
		},
		IsAnonymous:     false,
		Type:            "quiz",
		CorrectOptionID: 0,
	}
	if _, err := bot.Send(poll); err != nil {
		panic(err)
	}
}
