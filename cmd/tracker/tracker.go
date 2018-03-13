package main

import (
  "context"
  "fmt"
  "github.com/ozankasikci/apollo-telegram-tracker/firebase"
  "github.com/ozankasikci/apollo-telegram-tracker/telegrambot"
  "google.golang.org/api/iterator"
  "log"

)

func main() {
  ctx := context.Background()
  go telegrambot.InitTelegramBot()

  db := firebase.NewFirestoreClient(ctx)
  defer db.Close()

  iter := db.Collection("users").Documents(ctx)
  for {
    doc, err := iter.Next()
    if err == iterator.Done {
      break

    }

    if err != nil {
      log.Fatalf("failed to iterate %v", err)

    }
    fmt.Println(doc.Data())

  }

  done := make(chan bool)
  <- done
}
