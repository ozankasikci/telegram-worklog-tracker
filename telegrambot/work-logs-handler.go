package telegrambot

import (
	"context"
	"fmt"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
)

func NewWorkLogsHandler() *Handler {
	handler := &Handler{
		Description: "hours handler",
		Route:       "/hours",
	}

	function := func(ctx context.Context, m *tb.Message) {
		db := firebase.GetFirestoreClient(ctx)
		workLogsSnapshot, _ := db.Collection("work_logs").
			Where("user_id", "==", m.Sender.ID).
			Limit(20).
			Documents(ctx).
			GetAll()

		response := ""

		for i := 0; i < len(workLogsSnapshot); i++ {
			data := workLogsSnapshot[i].Data()
			response = response + fmt.Sprintf("check in time: %v\n", data["checkin_time"])

			if data["checkout_time"] != nil && data["checkout_time"] != "" {
				response = response + fmt.Sprintf("check out time: %v\n", data["checkout_time"])
			}
		}
		handler.ResponseMessage = response
	}

	handler.Func = function
	return handler
}
