package telegrambot

import (
	"context"
	tb "gopkg.in/tucnak/telebot.v2"
)

func PongHandlerFunction(ctx context.Context, h *Handler, m *tb.Message) {
	activityManager := GetActivityManager()
	activityManager.GetUserHashField(m.Sender.ID, "lastPongDate")

	if activityManager.GetUserHashField(m.Sender.ID, "lastPingDate") != "" {
		activityManager.CacheLastPongDate(m.Sender.ID)
		activityManager.DelUserHashField(m.Sender.ID, "lastPingDate")

		if h != nil {
			h.SetResponseMessage("Thanks for being here. A work log has been created!")
		}
	} else if h != nil {
		h.SetResponseMessage("Thanks but you were not pinged yet!")
	}

}

func NewPongHandler() *Handler {
	return &Handler{
		Description: "Pong handler",
		Route:       "/here",
		Func:        PongHandlerFunction,
	}
}
