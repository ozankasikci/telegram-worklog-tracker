package telegrambot

import (
	"context"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"sync"
	"time"
	"fmt"
)

var bot *tb.Bot
var once sync.Once

func FindOrCreateUser(ctx context.Context, m *tb.Message) {
	db := firebase.GetFirestoreClient(ctx)
	userSnapshot, err := db.Collection("users").
		Where("id", "==", m.Sender.ID).
		Documents(ctx).
		Next()

	if err != nil {
		fmt.Println("%v", err)
	}

	if userSnapshot == nil {
		fmt.Println("Creating user on firebase, user id:%d", m.Sender.ID)
		db.Collection("users").Add(ctx, map[string]interface{}{
			"id":       m.Sender.ID,
			"username": m.Sender.Username,
		})
	}
}

func RegisterHandlers(ctx context.Context, b *tb.Bot) {
	handlers := GetHandlers()

	for i := 0; i < len(handlers); i++ {
		handler := handlers[i]

		b.Handle(handler.Route, func(m *tb.Message) {
			// before each request, ensure user saved
			FindOrCreateUser(ctx, m)

			// run handler logic
			handler.Func(ctx, handler, m)

			response := handler.Description
			if handler.ResponseMessage != "" {
				response = handler.ResponseMessage
			}

			options := &tb.SendOptions{
				ParseMode: tb.ModeMarkdown,
			}
			b.Send(m.Sender, response, options)
		})
	}
}

func GetTelegramBot() (*tb.Bot, error) {
	var err error

	once.Do(func() {
		token := os.Getenv("TELEGRAM_TOKEN")
		bot, err = tb.NewBot(tb.Settings{
			Token:  token,
			Poller: &tb.LongPoller{Timeout: 60 * time.Second},
		})
	})

	return bot, err
}

func InitTelegramBot(ctx context.Context) {
	b, err := GetTelegramBot()
	RegisterHandlers(ctx, b)

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Start()
}
