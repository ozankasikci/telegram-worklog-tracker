package telegrambot

import (
	"context"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Handler struct {
	Description     string
	Route           string
	Func            func(context.Context, *tb.Message)
	ResponseMessage string
}

func GetHandlers() []*Handler {
	var handlers []*Handler

	handlers = append(
		handlers,
		NewCheckinHandler(),
		NewCheckoutHandler(),
		NewWorkLogsHandler(),
	)

	return handlers
}
