package services

import (
	"fmt"
	"github.com/Jeerihon/quizpool/keyboards"
	"github.com/Jeerihon/quizpool/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetQuizEvents(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	quizEvents := []models.QuizEvent{
		{
			QuizNumber:  1,
			Title:       "Quiz 1",
			Date:        "2021-10-01",
			Time:        "19:00",
			Link:        "https://www.google.com",
			Description: "This is the first quiz event",
		},
		{
			QuizNumber:  2,
			Title:       "Quiz 2",
			Date:        "2021-10-02",
			Time:        "19:00",
			Link:        "https://www.google.com",
			Description: "This is the second quiz event",
		},
	}

	for _, quizEvent := range quizEvents {
		messageText := fmt.Sprintf(
			"Quiz Number: %d\nTitle: %s\nDate: %s\nTime: %s\nLink: %s\nDescription: %s",
			quizEvent.QuizNumber,
			quizEvent.Title,
			quizEvent.Date,
			quizEvent.Time,
			quizEvent.Link,
			quizEvent.Description,
		)
		keyboard := keyboards.CreateAttendanceButton(quizEvent.QuizNumber)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
		msg.ReplyMarkup = keyboard

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
