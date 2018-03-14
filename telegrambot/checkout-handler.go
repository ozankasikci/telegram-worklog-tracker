package telegrambot

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

func NewCheckoutHandler() *Handler {
	function := func(ctx context.Context, h *Handler, m *tb.Message) {
		db := firebase.GetFirestoreClient(ctx)

		lastWorkLogSP, err := db.Collection("work_logs").
			OrderBy("checkin_time", firestore.Desc).
			Where("checkout_time", "==", "").
			Where("user_id", "==", m.Sender.ID).
			Limit(1).
			Documents(ctx).
			GetAll()

		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		if len(lastWorkLogSP) == 0 {
			println("No work log found")
			return
		}

		_, err = lastWorkLogSP[0].Ref.Update(ctx, []firestore.Update{
			{
				Path:  "checkout_time",
				Value: time.Now(),
			},
		})

		if err != nil {
			println(err)
		}
	}

	return &Handler{
		Description: "Check out handler",
		Route:       "/checkout",
		Func:        function,
	}
}
