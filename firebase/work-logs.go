package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	"time"
)

type WorkLogsOptions struct {
	Limit  int
	UserID int
	Minutes int
}

func FetchWorkLogs(ctx context.Context, options *WorkLogsOptions) ([]*firestore.DocumentSnapshot, error) {
	db := GetFirestoreClient(ctx)
	return db.Collection("work_logs").
		Where("user_id", "==", options.UserID).
		Limit(options.Limit).
		Documents(ctx).
		GetAll()
}

func CreateWorkLog(ctx context.Context, options *WorkLogsOptions) error {
	db := GetFirestoreClient(ctx)
	_, _, err := db.Collection("work_logs").Add(ctx, map[string]interface{}{
		"created_at": time.Now().Format(time.RFC3339),
		"user_id":    options.UserID,
		"minutes":    int(options.Minutes),
	})
	return err
}
