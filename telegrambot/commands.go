package telegrambot

import (
	"context"
	tb "gopkg.in/tucnak/telebot.v2"
)

func NewCommandsHandler() *Handler {
	function := func(ctx context.Context, h *Handler, m *tb.Message) {
		h.SetResponseMessage("Available Commands checkin, checkout, balance, work_log")
		h.SetResponseMessage("Available Commands  checkin, checkout, balance, worklog")
	}

	return &Handler{
		Description: "Commands handler",
		Route:       "/commands",
		Func:        function,
	}
}
