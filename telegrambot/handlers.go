package telegrambot

import (
	"context"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Handler struct {
	Description     string
	Route           string
	Func            func(context.Context, *Handler, *tb.Message)
	ResponseMessage string
}

func (h *Handler) SetReponseMessage(m string) {
	h.ResponseMessage = m
}

func GetHandlers() []*Handler {
	var handlers []*Handler

	handlers = append(
		handlers,
		NewCheckinHandler(),
		NewCheckoutHandler(),
		NewWorkLogsHandler(),
		NewBalanceHandler(),
	)

	return handlers
}
