package main

import (
	"context"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	"github.com/ozankasikci/apollo-telegram-tracker/telegrambot"
)

func main() {
	ctx := context.Background()
	go telegrambot.InitTelegramBot(ctx)

	db := firebase.GetFirestoreClient(ctx)
	defer db.Close()

	telegrambot.GetActivityManager()

	done := make(chan bool)
	<-done
}
