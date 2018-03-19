package telegrambot

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	ActiveTimeout    = 1
	WorklogThreshold = 1
	LoopTime         = 1
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

func (am *ActivityManager) Init() {
	task := func() {
		activeUsers, err := am.redis.SMembers("active_users").Result()
		fmt.Println("Active users: %v", activeUsers)
		if err != nil {
			log.Fatalln(err)
		}

		for i := 0; i < len(activeUsers); i++ {
			userId := activeUsers[i]
			bot, _ := GetTelegramBot()
			id, _ := strconv.Atoi(userId)

			user := &telebot.User{ID: id}
			userHash := am.GetUserHashAll(id)

			lastCheckinDate, _ := time.Parse(time.RFC3339, userHash["lastCheckinDate"])
			lastPingDate, _ := time.Parse(time.RFC3339, userHash["lastPingDate"])
			lastPongDate, _ := time.Parse(time.RFC3339, userHash["lastPongDate"])

			// time is up, clear user cache
			if lastPingDate.IsZero() == false && time.Since(lastPingDate).Minutes() >= ActiveTimeout {
				am.RemoveFromActiveUsers(id)
				bot.Send(user, "Don't forget to check in if you haven't already!")
			} else if lastPingDate.IsZero() &&
				time.Since(lastCheckinDate).Minutes() > WorklogThreshold &&
				(lastPongDate.IsZero() || time.Since(lastPongDate).Minutes() >= WorklogThreshold) {

				bot.Send(user, "Are you still here? Answer /here")
				am.CacheLastPingDate(id)
			}
		}
	}

	// continuously check if users are active
	ticker := time.NewTicker(LoopTime * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <- ticker.C:
				task()
			case <- quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (am *ActivityManager) AddToActiveUsers(userId int) {
	fmt.Println("adding to active users user")
	am.redis.SAdd("active_users", userId)
}

func (am *ActivityManager) RemoveFromActiveUsers(userId int) {
	am.redis.SRem("active_users", userId)
	am.ClearUserHash(userId)
}

func (am *ActivityManager) GetUserHashField(userId int, field string) string {
	res, _ := am.redis.HGet(GetUserKey(userId), field).Result()
	return res
}

func (am *ActivityManager) DelUserHashField(userId int, field string) {
	am.redis.HDel(GetUserKey(userId), field)
}

func (am *ActivityManager) GetUserHashAll(userId int) map[string]string {
	res, _ := am.redis.HGetAll(GetUserKey(userId)).Result()
	return res
}

func (am *ActivityManager) CacheLastPingDate(userId int) {
	am.redis.HSet(GetUserKey(userId), "lastPingDate", time.Now().Format(time.RFC3339))
}

func (am *ActivityManager) CacheLastPongDate(userId int) {
	am.redis.HSet(GetUserKey(userId), "lastPongDate", time.Now().Format(time.RFC3339))
}

func (am *ActivityManager) CacheLastCheckinDate(userId int) {
	am.redis.HSetNX(GetUserKey(userId), "lastCheckinDate", time.Now().Format(time.RFC3339))
}

func (am *ActivityManager) ClearUserHash(userId int) {
	am.redis.Del(GetUserKey(userId))
}
