package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"os"
	"sync"
)

var firestoreClient *firestore.Client
var once sync.Once

// singleton firestore client
func GetFirestoreClient(ctx context.Context) *firestore.Client {
	once.Do(func() {
		firestoreClient = newFirestoreClient(ctx)
	})
	return firestoreClient
}

func NewFirebaseAdminApp(ctx context.Context) *firebase.App {
	pwd, _ := os.Getwd()
	conf := &firebase.Config{
		DatabaseURL: "https://apollo-telegram-bot.firebaseio.com",
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(pwd + "/firebase/service-account.json")

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	return app
}

func newFirestoreClient(ctx context.Context) *firestore.Client {
	firebaseAdmin := NewFirebaseAdminApp(ctx)

	db, err := firebaseAdmin.Firestore(ctx)

	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	return db
}
