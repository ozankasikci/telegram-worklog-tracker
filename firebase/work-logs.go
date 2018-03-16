package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	"time"
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

func CreateWorkLog(ctx context.Context, options *WorkLogsOptions) error {
	db := GetFirestoreClient(ctx)
	_, _, err := db.Collection("work_logs").Add(ctx, map[string]interface{}{
		"checkin_time":  time.Now(),
		"checkout_time": "",
		"user_id":       options.UserID,
	})
	return err
}

//func ValidateWorkLog(ctx context.Context, options *WorkLogsOptions) {
//	db := GetFirestoreClient(ctx)
//	return db.Collection("work_logs").
//		Where("user_id", "==", options.UserID).
//		Limit(options.Limit).
//		Documents(ctx).
//		GetAll()
//}