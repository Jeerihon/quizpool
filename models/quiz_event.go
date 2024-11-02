package models

type QuizEvent struct {
	QuizNumber  int    `json:"quiz_number"`
	Date        string `json:"date"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Time        string `json:"time"`
}
