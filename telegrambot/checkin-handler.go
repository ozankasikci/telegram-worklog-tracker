package telegrambot

import (
	"context"
	tb "gopkg.in/tucnak/telebot.v2"
)

func CheckinHandlerFunction(ctx context.Context, h *Handler, m *tb.Message) {
	FindOrCreateUser(ctx, m)
	am := GetActivityManager()
	am.AddToActiveUsers(m.Sender.ID)
	am.CacheLastCheckinDate(m.Sender.ID)
	h.SetResponseMessage("Successfully checked in. (type /commands to see available commands)")
}

func NewCheckinHandler() *Handler {
	return &Handler{
		Description: "Check in handler",
		Route:       "/checkin",
		Func:        CheckinHandlerFunction,
	}
}
