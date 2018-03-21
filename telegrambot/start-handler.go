package telegrambot

import (
	"context"
	tb "gopkg.in/tucnak/telebot.v2"
)

func sendResponse(handler *Handler, b *tb.Bot, m *tb.Message) {
	response := handler.Description
	if handler.ResponseMessage != "" {
		response = handler.ResponseMessage
	}

	options := &tb.SendOptions{
		ParseMode: tb.ModeMarkdown,
	}
	b.Send(m.Sender, response, options)
}

func NewStartHandler() *Handler {
	function := func(ctx context.Context, h *Handler, m *tb.Message) {
		checkinBtn := tb.ReplyButton{Text: "Check in"}
		checkoutBtn := tb.ReplyButton{Text: "Check out"}
		earningsBtn := tb.ReplyButton{Text: "Earned DAEs"}
		workLogsBtn := tb.ReplyButton{Text: "Work Logs"}

		replyKeys := [][]tb.ReplyButton{
			[]tb.ReplyButton{checkinBtn, checkoutBtn},
			[]tb.ReplyButton{earningsBtn, workLogsBtn},
		}

		b, _ := GetTelegramBot()

		b.Handle(&checkinBtn, func(m *tb.Message) {
			CheckinHandlerFunction(ctx, h, m)
			sendResponse(h, b, m)
		})

		b.Handle(&checkoutBtn, func(m *tb.Message) {
			CheckoutHandlerFunction(ctx, h, m)
			sendResponse(h, b, m)
		})

		b.Handle(&earningsBtn, func(m *tb.Message) {
			BalanceHandlerFunction(ctx, h, m)
			sendResponse(h, b, m)
		})

		b.Handle(&workLogsBtn, func(m *tb.Message) {
			WorkLogsHandlerFunction(ctx, h, m)
			sendResponse(h, b, m)
		})

		b.Send(m.Sender, "Welcome!", &tb.ReplyMarkup{
			ReplyKeyboard:  replyKeys,
		})
	}

	return &Handler{
		Description: "",
		Route:       "/start",
		Func:        function,
	}
}
