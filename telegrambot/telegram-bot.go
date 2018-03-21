package telegrambot

import (
	"context"
	"fmt"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"sync"
	"time"
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
			"id":         m.Sender.ID,
			"username":   m.Sender.Username,
			"created_at": time.Now().Format(time.RFC3339),
		})
	}
}

func RegisterHandlers(ctx context.Context, b *tb.Bot) {
	handlers := GetHandlers()

	for i := 0; i < len(handlers); i++ {
		handler := handlers[i]

		b.Handle(handler.Route, func(m *tb.Message) {
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

	if os.Getenv("TELEGRAM_TOKEN") == "" {
		log.Fatalln("Please set TELEGRAM_TOKEN env var")
	}

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

	if err != nil {
		log.Fatal(err)
		return
	}

	RegisterHandlers(ctx, b)

	b.Start()
}
