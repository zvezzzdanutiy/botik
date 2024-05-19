package TelegramBot

import (
	"context"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ChatInfo struct {
	stop   chan bool
	ticker *time.Ticker
}

var chats = make(map[int64]*ChatInfo)

func (d *Domain) StartBot() {
	d.botApi.Debug = true

	log.Printf("Authorized on account %s", d.botApi.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	updates, _ := d.botApi.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == "/start" {

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать! Выберите действие:")
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Этнический юмор"),
						tgbotapi.NewKeyboardButton("Профессиональный юмор"),
						tgbotapi.NewKeyboardButton("Детский юмор"),
						tgbotapi.NewKeyboardButton("Подписаться"),
					),
				)
				d.botApi.Send(msg)

			} else if update.Message.Text == "Этнический юмор" || update.Message.Text == "Ржать еще!" {

				anekdot, err := d.anekdotProvider.GetJoke(context.Background(), "1")
				if err != nil {
					log.Println("Ошибка GetAnekdot")
				}
				msg := tgbotapi.NewMessage(int64(update.Message.From.ID), anekdot)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Ржать еще!"),
						tgbotapi.NewKeyboardButton("Сменить категорию"),
					),
				)
				d.botApi.Send(msg)

			} else if update.Message.Text == "Профессиональный юмор" || update.Message.Text == "Ржать еще!!" {

				anekdot, err := d.anekdotProvider.GetJoke(context.Background(), "2")
				if err != nil {
					log.Println("Ошибка GetAnekdot")
				}
				msg := tgbotapi.NewMessage(int64(update.Message.From.ID), anekdot)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Ржать еще!!"),
						tgbotapi.NewKeyboardButton("Сменить категорию"),
					),
				)
				d.botApi.Send(msg)

			} else if update.Message.Text == "Детский юмор" || update.Message.Text == "Ржать еще!!!" {

				anekdot, err := d.anekdotProvider.GetJoke(context.Background(), "3")
				if err != nil {
					log.Println("Ошибка GetAnekdot")
				}
				msg := tgbotapi.NewMessage(int64(update.Message.From.ID), anekdot)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Ржать еще!!!"),
						tgbotapi.NewKeyboardButton("Сменить категорию"),
					),
				)
				d.botApi.Send(msg)

			} else if update.Message.Text == "Подписаться" {

				stop := make(chan bool)
				ticker := time.NewTicker(5 * time.Second)
				chats[update.Message.Chat.ID] = &ChatInfo{
					stop:   stop,
					ticker: ticker,
				}
				go d.autoSendJokes(update.Message.Chat.ID, stop)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы подписались на рассылку анекдотов.")
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Этнический юмор"),
						tgbotapi.NewKeyboardButton("Профессиональный юмор"),
						tgbotapi.NewKeyboardButton("Детский юмор"),
						tgbotapi.NewKeyboardButton("Отписаться"),
					),
				)
				d.botApi.Send(msg)

			} else if update.Message.Text == "Отписаться" {

				if chatInfo, ok := chats[update.Message.Chat.ID]; ok {
					chatInfo.stop <- true
					chatInfo.ticker.Stop()
					delete(chats, update.Message.Chat.ID)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы отписались от рассылки анекдотов.")
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Этнический юмор"),
						tgbotapi.NewKeyboardButton("Профессиональный юмор"),
						tgbotapi.NewKeyboardButton("Детский юмор"),
						tgbotapi.NewKeyboardButton("Подписаться"),
					),
				)
				d.botApi.Send(msg)

			} else if update.Message.Text == "Сменить категорию" {

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите категорию анекдота:")
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Этнический юмор"),
						tgbotapi.NewKeyboardButton("Профессиональный юмор"),
						tgbotapi.NewKeyboardButton("Детский юмор"),
						tgbotapi.NewKeyboardButton("Подписаться"),
					),
				)
				d.botApi.Send(msg)

			}
		}
	}
}
func (d *Domain) autoSendJokes(chatID int64, stop chan bool) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-stop:
			return
		case <-ticker.C:
			anekdot, err := d.anekdotProvider.GetJoke(context.Background(), "")
			if err != nil {
				log.Println("Ошибка GetAnekdot")
				continue
			}
			msg := tgbotapi.NewMessage(chatID, anekdot)
			d.botApi.Send(msg)
		}
	}
}
