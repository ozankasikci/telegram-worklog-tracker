package telegrambot

import (
	"context"
	tb "gopkg.in/tucnak/telebot.v2"
)

func CheckinHandlerFunction(ctx context.Context, h *Handler, m *tb.Message) {
	FindOrCreateUser(ctx, m)
	am := GetActivityManager()
	am.AddToActiveUsers(m.Sender.ID)

	// sets only if last check in date doesn't exist
	created := am.CacheLastCheckinDate(m.Sender.ID)

	if created == true {
		h.SetResponseMessage("Successfully checked in!")
	} else {
		h.SetResponseMessage("You are already checked in!")
	}
}

func NewCheckinHandler() *Handler {
	return &Handler{
		Description: "Check in handler",
		Route:       "/checkin",
		Func:        CheckinHandlerFunction,
	}
}
