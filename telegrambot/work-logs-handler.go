package telegrambot

import (
	"context"
	"fmt"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
)

var workhourSnippet = `
*checkin:* %s
*checkout:* %s
`

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
			response = response + fmt.Sprintf(workhourSnippet, data["checkin_time"], data["checkout_time"])
		}

		handler.ResponseMessage = response
	}

	handler.Func = function
	return handler
}
