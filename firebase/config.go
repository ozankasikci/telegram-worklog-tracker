package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
)

func FetchConfig(ctx context.Context) (*firestore.DocumentSnapshot, error) {
	db := GetFirestoreClient(ctx)
	return db.Collection("config").
		Limit(1).
		Documents(ctx).
		Next()
}
