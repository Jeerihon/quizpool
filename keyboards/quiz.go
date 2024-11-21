package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateAttendanceButton(quizEventID string) tgbotapi.InlineKeyboardMarkup {
	callbackData := "create_attendance_poll:" + quizEventID
	button := tgbotapi.NewInlineKeyboardButtonData("Создать опрос", callbackData)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button))
	return keyboard
}
