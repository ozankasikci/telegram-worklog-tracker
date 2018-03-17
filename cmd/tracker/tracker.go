package main

import (
	"context"
	"github.com/ozankasikci/apollo-telegram-tracker/telegrambot"
)

func main() {
	ctx := context.Background()

	activityManager := telegrambot.GetActivityManager()
	activityManager.Init()

	telegrambot.InitTelegramBot(ctx)
}
