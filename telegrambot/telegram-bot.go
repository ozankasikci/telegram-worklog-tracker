package telegrambot

import (
	"context"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	"strconv"
)

func FindOrCreateUser(ctx context.Context, m *tb.Message) {
	db := firebase.GetFirestoreClient(ctx)
	userSnapshot, err := db.Doc(strconv.Itoa(m.Sender.ID)).Get(ctx)

	if err != nil {
		println(err)
	}

	if userSnapshot == nil {
		db.Collection("users").Add(ctx, map[string]interface{}{
			"id": m.Sender.ID,
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
			handler.Func(ctx, m)

			response := handler.Description
			if handler.ResponseMessage != "" {
				response = handler.ResponseMessage
			}

			b.Send(m.Sender, response)
		})
	}
}

func InitTelegramBot(ctx context.Context) {
	b, err := tb.NewBot(tb.Settings{
		Token:  "492394305:AAEGTfMTO2qxqqa7BGkuhvzLgnHeA7Ek7C4",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	RegisterHandlers(ctx, b)

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Start()
}
