package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
)

type WorkLogsOptions struct {
	Limit  int
	UserID int
}

func FetchWorkLogs(ctx context.Context, options *WorkLogsOptions) ([]*firestore.DocumentSnapshot, error) {
	db := GetFirestoreClient(ctx)
	return db.Collection("work_logs").
		Where("user_id", "==", options.UserID).
		Limit(options.Limit).
		Documents(ctx).
		GetAll()
}
