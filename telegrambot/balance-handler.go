package telegrambot

import (
	"context"
	"fmt"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

var paymentsSnippet = `
*Total Minutes:* %d
*Total DAE Earned:* %d
`

func BalanceHandlerFunction(ctx context.Context, h *Handler, m *tb.Message) {
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
	paymentPerMinute, ok := config["payment_per_minute"].(int64)
	if !ok {
		log.Printf("got data of type %T but wanted int64", paymentPerMinute)
	}

	var totalMinutes int64
	totalMinutes = 0
	for i := 0; i < len(workLogsSnapshots); i++ {
		snapshot := workLogsSnapshots[i]
		data := snapshot.Data()
		minutesInt, ok := data["minutes"].(int64)
		if !ok {
			log.Printf("got data of type %T but wanted int64", minutesInt)
		}

		totalMinutes += minutesInt
	}

	h.SetResponseMessage(fmt.Sprintf(paymentsSnippet, totalMinutes, totalMinutes*paymentPerMinute))
	fmt.Println(h.ResponseMessage)
}

func NewBalanceHandler() *Handler {
	return &Handler{
		Description: "Balance handler",
		Route:       "/balance",
		Func:        BalanceHandlerFunction,
	}
}
