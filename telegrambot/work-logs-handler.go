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
	function := func(ctx context.Context, h *Handler, m *tb.Message) {
		response := ""
		options := &firebase.WorkLogsOptions{Limit: 50, UserID: m.Sender.ID}
		workLogsSnapshot, err := firebase.FetchWorkLogs(ctx, options)

		if err != nil {
			println(err)
			return
		}

		for i := 0; i < len(workLogsSnapshot); i++ {
			data := workLogsSnapshot[i].Data()
			response = response + fmt.Sprintf(workhourSnippet, data["checkin_time"], data["checkout_time"])
		}

		h.SetResponseMessage(response)
	}

	return &Handler{
		Description: "Work Log handler",
		Route:       "/work_log",
		Func:        function,
	}

}
