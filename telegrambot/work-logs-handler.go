package telegrambot

import (
	"context"
	"fmt"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
)

var workhourSnippet = `
*Work Log Created at:* %s
`

func WorkLogsHandlerFunction(ctx context.Context, h *Handler, m *tb.Message) {
	response := ""
	options := &firebase.WorkLogsOptions{Limit: 50, UserID: m.Sender.ID}
	workLogsSnapshot, err := firebase.FetchWorkLogs(ctx, options)

	if err != nil {
		println(err)
		return
	}

	for i := 0; i < len(workLogsSnapshot); i++ {
		data := workLogsSnapshot[i].Data()
		response = response + fmt.Sprintf(workhourSnippet, data["created_at"])
	}

	h.SetResponseMessage(response)
}

func NewWorkLogsHandler() *Handler {
	return &Handler{
		Description: "No work logs yet!",
		Route:       "/worklog",
		Func:        WorkLogsHandlerFunction,
	}

}
