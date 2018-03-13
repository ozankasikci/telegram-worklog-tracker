package telegrambot

import (
	"context"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

func NewCheckinHandler() *Handler {
	function := func(ctx context.Context, m *tb.Message) {
		db := firebase.GetFirestoreClient(ctx)
		db.Collection("work_logs").Add(ctx, map[string]interface{}{
			"checkin_time":  time.Now(),
			"checkout_time": "",
			"user_id":       m.Sender.ID,
		})
	}

	return &Handler{
		Description: "Check in handler",
		Route:       "/checkin",
		Func:        function,
	}
}
