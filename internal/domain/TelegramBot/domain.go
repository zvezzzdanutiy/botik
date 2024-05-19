package TelegramBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Domain struct {
	botApi          *tgbotapi.BotAPI
	anekdotProvider AnekdotProvider
}

func New(token string, anekdotProvider AnekdotProvider) (*Domain, error) {
	s, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Domain{
		botApi:          s,
		anekdotProvider: anekdotProvider,
	}, nil
}
