package main

import (
	"context"
	"github.com/jasonlvhit/gocron"
	"github.com/ozankasikci/apollo-telegram-tracker/telegrambot"
)

func main() {
	ctx := context.Background()

	activityManager := telegrambot.GetActivityManager()
	activityManager.Init()

	<-gocron.Start()

	telegrambot.InitTelegramBot(ctx)
}
