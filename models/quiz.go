package models

type QuizEvent struct {
	Title      string `json:"title"`
	QuizNumber string `json:"quiz_number"`
	Place      string `json:"place"`
	Info       string `json:"info"`
	Link       string `json:"link"`
	Date       string `json:"date"`
	Time       string `json:"time"`
}
