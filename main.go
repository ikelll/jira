package main

import (
	"./bot"
	"log"
)

func main() {
	botToken := ""

	myBot, err := bot.NewBot(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.HandleMessages(myBot)
}
