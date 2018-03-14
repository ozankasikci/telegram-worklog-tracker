package telegrambot

import (
	"context"
	"fmt"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
	"strconv"
)

const (
	CheckedIn = iota
	CheckedOut
	ClockecOut
)

func NewCheckinHandler() *Handler {
	function := func(ctx context.Context, h *Handler, m *tb.Message) {

		db := firebase.GetFirestoreClient(ctx)
		_, _, err := db.Collection("work_logs").Add(ctx, map[string]interface{}{
			"checkin_time":  time.Now(),
			"checkout_time": "",
			"user_id":       m.Sender.ID,
		})

		if err != nil {
			h.SetReponseMessage(fmt.Sprintf("Failed to checkin, err:%v", err))
		}

		GetActivityManager().Set(strconv.Itoa(m.Sender.ID), CheckedIn, 0)
		h.SetReponseMessage("Successfully checked in.")
	}

	return &Handler{
		Description: "Check in handler",
		Route:       "/checkin",
		Func:        function,
	}
}
