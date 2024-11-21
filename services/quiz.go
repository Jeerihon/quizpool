package services

import (
	"fmt"
	"github.com/Jeerihon/quizpool/keyboards"
	"github.com/Jeerihon/quizpool/models"
	"github.com/Jeerihon/quizpool/utils"
	"github.com/PuerkitoBio/goquery"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"strings"
)

func GetSchedule(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	url := "https://astana.quizplease.ru"
	events := ParseScheduleEvents(url+"/schedule", ".schedule-block")

	for _, quizEvent := range events {
		messageText := fmt.Sprintf(
			"[%s](%s) *%s*\n%s %s\n%s\n%s",
			utils.EscapeMarkdownV2(quizEvent.Title),
			utils.EscapeMarkdownV2(url+quizEvent.Link),
			utils.EscapeMarkdownV2(quizEvent.QuizNumber),
			utils.EscapeMarkdownV2(quizEvent.Date),
			utils.EscapeMarkdownV2(quizEvent.Time),
			utils.EscapeMarkdownV2(quizEvent.Place),
			utils.EscapeMarkdownV2(quizEvent.Info),
		)
		keyboard := keyboards.CreateAttendanceButton(quizEvent.QuizNumber)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
		msg.ReplyMarkup = keyboard
		msg.ParseMode = "markdownv2"

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}

func CreateAttendancePoll(bot *tgbotapi.BotAPI, update tgbotapi.Update, quizEventId string) {
	poll := tgbotapi.SendPollConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           update.CallbackQuery.Message.Chat.ID,
			ReplyToMessageID: update.CallbackQuery.Message.MessageID,
		},
		Question: "Вы идете на игру?",
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

func ParseScheduleEvents(url string, selector string) []models.QuizEvent {
	res, err := http.Get(url)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	events := []models.QuizEvent{}

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		events = append(events, models.QuizEvent{
			Title:      strings.TrimSpace(strings.Replace(s.Find(".h2").First().Text(), "ASTANA", "", -1)),
			QuizNumber: s.Find(".h2").Eq(1).Text(),
			Place:      s.Find(".schedule-block-info-bar").Text(),
			Info:       s.Find(".schedule-days-left-text").Text(),
			Link:       s.Find("a").AttrOr("href", ""),
			Date:       strings.TrimSpace(s.Find(".h3").Text()),
			Time:       s.Find(".schedule-info:nth-child(2) .techtext").Text(),
		})
	})

	return events
}
