package telegrambot

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

func RegisterHandlers(b *tb.Bot) {
	handlers := GetHandlers()

	for i := 0; i < len(handlers); i++ {
		handler := handlers[i]

		b.Handle(handler.Route, func(m *tb.Message) {
			b.Send(m.Sender, handler.Description)
		})
	}
}

func InitTelegramBot() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "492394305:AAEGTfMTO2qxqqa7BGkuhvzLgnHeA7Ek7C4",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	RegisterHandlers(b)

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Start()
}
