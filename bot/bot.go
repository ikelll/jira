package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type MyBot struct {
	*tgbotapi.BotAPI
}

func NewBot(token string) (*MyBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &MyBot{bot}, nil
}
