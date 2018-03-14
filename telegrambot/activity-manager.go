package telegrambot

import (
	"github.com/go-redis/redis"
	"log"
	"sync"
	"os"
)

var client *redis.Client
var redisOnce sync.Once

func GetActivityManager() *redis.Client {
	redisOnce.Do(func() {
		redisHost := "127.0.0.1"
		if os.Getenv("REDIS_HOST") != "" {
			redisHost = os.Getenv("REDIS_HOST")
		}
		
		client = redis.NewClient(&redis.Options{
			Addr:     redisHost + ":6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalln("Can not reach redis")
	}

	return client
}

