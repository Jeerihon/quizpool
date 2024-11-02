package main

import (
	"github.com/Jeerihon/quizpool/client"
	"github.com/Jeerihon/quizpool/handlers"
)

func main() {
	bot := client.Init()

	handlers.InitCommands(bot)
	handlers.Init(bot)
}
