package telegrambot

import (
	"github.com/go-redis/redis"
	"log"
	"sync"
	"os"
)

var activityManager *ActivityManager
var redisOnce sync.Once

type ActivityManager struct {
	redis *redis.Client
}

func GetActivityManager() *ActivityManager {
	redisOnce.Do(func() {
		redisHost := "127.0.0.1"
		if os.Getenv("REDIS_HOST") != "" {
			redisHost = os.Getenv("REDIS_HOST")
		}
		
		redis := redis.NewClient(&redis.Options{
			Addr:     redisHost + ":6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		activityManager = &ActivityManager{
			redis: redis,
		}
	})

	_, err := activityManager.redis.Ping().Result()
	if err != nil {
		log.Fatalln("Can not reach redis")
	}

	return activityManager
}

