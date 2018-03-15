package main

import (
	"context"
	"github.com/ozankasikci/apollo-telegram-tracker/firebase"
	"github.com/ozankasikci/apollo-telegram-tracker/telegrambot"
	"github.com/jasonlvhit/gocron"
)

func main() {
	ctx := context.Background()
	go telegrambot.InitTelegramBot(ctx)

	db := firebase.GetFirestoreClient(ctx)
	defer db.Close()

	activityManager := telegrambot.GetActivityManager()
	activityManager.Init()

	<- gocron.Start()

	done := make(chan bool)
	<-done
}
