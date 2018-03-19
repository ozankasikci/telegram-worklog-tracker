package telegrambot

import (
	"context"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	CheckedIn = iota
	CheckedOut
	ClockecOut
)

func NewCheckinHandler() *Handler {
	function := func(ctx context.Context, h *Handler, m *tb.Message) {
		FindOrCreateUser(ctx, m)
		activityManager := GetActivityManager()
		activityManager.AddToActiveUsers(m.Sender.ID)
		activityManager.CacheLastCheckinDate(m.Sender.ID)
		h.SetResponseMessage("Successfully checked in. (type /commands to see available commands)")
	}

	return &Handler{
		Description: "Check in handler",
		Route:       "/checkin",
		Func:        function,
	}
}
