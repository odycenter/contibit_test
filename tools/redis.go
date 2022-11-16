package tools

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var mutex sync.Mutex

func RedisClient(add string, port int, passwd string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     add + ":" + strconv.Itoa(port),
		Password: passwd, // no password set
		DB:       db,     // use default DB
	})
	return client
}

func RedisLock(redis *redis.Client, key string, value string, timeout time.Duration) bool {
	mutex.Lock()
	defer mutex.Unlock()
	temp, err := redis.SetNX(key, value, timeout).Result()
	if err != nil {
		log.Println("err", "Redis Lock error", key+":("+value+")", err.Error())
	}
	return temp
}

func RedisUnLock(redis *redis.Client, key string, value string) bool {
	temp, _ := redis.Get(key).Result()
	if temp == value {
		nums, err := redis.Del(key).Result()
		if err != nil {
			log.Println("err", "Redis UnLock error", key+":("+value+")", "")
		}
		if nums == 1 {
			return true
		}
	}
	return false
}
