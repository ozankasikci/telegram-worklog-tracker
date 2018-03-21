package telegrambot

import (
	"context"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

func CheckoutHandlerFunction(ctx context.Context, h *Handler, m *tb.Message) {
	activityManager := GetActivityManager()
	activityManager.RemoveFromActiveUsers(m.Sender.ID)
	fmt.Println("Checking out user, %d", m.Sender.ID)
	h.SetResponseMessage("Successfully checked out.")
}

func NewCheckoutHandler() *Handler {
	return &Handler{
		Description: "Check out handler",
		Route:       "/checkout",
		Func:        CheckoutHandlerFunction,
	}
}
