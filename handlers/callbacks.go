package handlers

import (
	"github.com/Jeerihon/quizpool/services"
	"github.com/Jeerihon/quizpool/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmd, quizEventId := utils.GetKeyValue(update.CallbackQuery.Data)

	switch {
	case cmd == "create_attendance_poll":
		services.CreateAttendancePoll(bot, update, quizEventId)
	}
}
