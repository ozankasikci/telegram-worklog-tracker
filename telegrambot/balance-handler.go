package telegrambot

import (
	"context"
	"fmt"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

var paymentsSnippet = `
*Total Hours:* %d
*Total DAE Earned:* %d
`

func NewBalanceHandler() *Handler {
	function := func(ctx context.Context, h *Handler, m *tb.Message) {
		configSnapshot, err := firebase.FetchConfig(ctx)
		if err != nil {
			log.Fatalln(err)
		}

		options := &firebase.WorkLogsOptions{Limit: 50, UserID: m.Sender.ID}
		workLogsSnapshots, err := firebase.FetchWorkLogs(ctx, options)
		if err != nil {
			log.Fatalln(err)
		}

		config := configSnapshot.Data()
		hourlyPay, ok := config["hourly_pay"].(int64)
		if !ok {
			log.Printf("got data of type %T but wanted int64", hourlyPay)
		}

		var totalHours int64
		totalHours = 0
		for i := 0; i < len(workLogsSnapshots); i++ {
			totalHours += 1
		}

		h.SetResponseMessage(fmt.Sprintf(paymentsSnippet, totalHours, totalHours*hourlyPay))
	}

	return &Handler{
		Description: "Balance handler",
		Route:       "/balance",
		Func:        function,
	}
}
