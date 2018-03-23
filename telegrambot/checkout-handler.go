package telegrambot

import (
	"context"
	"fmt"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

func CheckoutHandlerFunction(ctx context.Context, h *Handler, m *tb.Message) {
	am := GetActivityManager()
	userHash := am.GetUserHashAll(m.Sender.ID)
	lastCheckinDate, _ := time.Parse(time.RFC3339, userHash["lastCheckinDate"])

	minutes := time.Since(lastCheckinDate).Minutes()
	minutesInt := int(minutes)
	wlo := &firebase.WorkLogsOptions{Minutes: minutesInt, UserID: m.Sender.ID}
	am.CreateWorkLog(ctx, wlo)

	am.RemoveFromActiveUsers(m.Sender.ID)
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
