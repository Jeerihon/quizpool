package constants

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var BotCommands = []tgbotapi.BotCommand{
	{Command: "schedule", Description: "Расписание игр на ближайшее время"},
	{Command: "start", Description: "Начать работу с ботом"},
	{Command: "help", Description: "Помощь"},
}
