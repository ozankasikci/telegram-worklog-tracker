package telegrambot

import (
	"context"
	tb "gopkg.in/tucnak/telebot.v2"
)

func NewCheckoutHandler() *Handler {
	function := func(ctx context.Context, h *Handler, m *tb.Message) {
		activityManager := GetActivityManager()
		activityManager.RemoveFromActiveUsers(m.Sender.ID)
		h.SetResponseMessage("Successfully checked out.")
	}

	return &Handler{
		Description: "Check out handler",
		Route:       "/checkout",
		Func:        function,
	}
}
