package telegrambot

import (
	"github.com/go-redis/redis"
	"log"
	"sync"
	"os"
	"fmt"
	"time"
	"strconv"
	"github.com/jasonlvhit/gocron"
	"gopkg.in/tucnak/telebot.v2"
)

var activityManager *ActivityManager
var redisOnce sync.Once

type ActivityManager struct {
	redis *redis.Client
}

type CacheUserOptions struct {
	lastCheckInDate time.Time
}

func GetUserKey(userId int) string {
	return "users:" + strconv.Itoa(userId)
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

func (am *ActivityManager) Init()  {
	task := func() {
		activeUsers, err := am.redis.SMembers("active_users").Result()
		if err != nil {
			log.Fatalln(err)
		}

		for i := 0; i < len(activeUsers); i++ {
			userId := activeUsers[i]
			bot, _ := GetTelegramBot()
			id, _ := strconv.Atoi(userId)

			user := &telebot.User{ ID: id }
			bot.Send(user, "Are you still here? Answer /here")
		}
	}

	gocron.Every(30).Minutes().Do(task)
}

func (am *ActivityManager) AddToActiveUsers(userId int)  {
	fmt.Println("adding to active users user")
	am.redis.SAdd("active_users", userId)
}

func (am *ActivityManager) RemoveFromActiveUsers(userId int)  {
	am.redis.SRem("active_users", userId)
	am.redis.Del(GetUserKey(userId))
}

func (am *ActivityManager) CacheUser(userId int) {
	fmt.Println("caching user")
	am.redis.HSetNX(GetUserKey(userId), "lastCheckinDate", time.Now().String())
}
