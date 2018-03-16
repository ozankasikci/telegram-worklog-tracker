package telegrambot

import (
	"context"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"strconv"
	"time"
	"sync"
)

var bot  *tb.Bot
var once sync.Once

func FindOrCreateUser(ctx context.Context, m *tb.Message) {
	db := firebase.GetFirestoreClient(ctx)
	userSnapshot, err := db.Doc(strconv.Itoa(m.Sender.ID)).Get(ctx)

	if err != nil {
		println(err)
	}

	if userSnapshot == nil {
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
			GetActivityManager().redis.SetNX("chat_id", m.Chat.ID, 0)

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

func GetTelegramBot() (*tb.Bot, error){
	var err error

	once.Do(func() {
		bot, err = tb.NewBot(tb.Settings{
			Token:  "492394305:AAEGTfMTO2qxqqa7BGkuhvzLgnHeA7Ek7C4",
			Poller: &tb.LongPoller{Timeout: 10 * time.Second},
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
