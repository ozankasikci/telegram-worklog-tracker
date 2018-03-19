package telegrambot

import (
	"context"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
)

func NewPongHandler() *Handler {
	function := func(ctx context.Context, h *Handler, m *tb.Message) {
		activityManager := GetActivityManager()
		activityManager.GetUserHashField(m.Sender.ID, "lastPongDate")

		if activityManager.GetUserHashField(m.Sender.ID, "lastPingDate") != "" {
			options := &firebase.WorkLogsOptions{UserID: m.Sender.ID}
			firebase.CreateWorkLog(ctx, options)
			activityManager.CacheLastPongDate(m.Sender.ID)
			activityManager.DelUserHashField(m.Sender.ID, "lastPingDate")
			h.SetResponseMessage("Thanks for being here. A work log has been created!")
		} else {
			h.SetResponseMessage("Thanks but you were not pinged yet!")
		}

	}

	return &Handler{
		Description: "Pong handler",
		Route:       "/here",
		Func:        function,
	}
}
