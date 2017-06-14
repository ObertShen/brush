package db

import (
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var redisClientIns struct {
	*redis.Client
	sync.Mutex
}

// GetRedisClientIns 获取redis.Client的单例
func GetRedisClientIns() *redis.Client {
	if redisClientIns.Client == nil {
		redisClientIns.Lock()
		if redisClientIns.Client == nil {
			redisClientIns.Client = NewClient()
		}

		redisClientIns.Unlock()
	}

	return redisClientIns.Client
}

// NewClient 创建 redis 连接
func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:        "go.sna.com:6379",
		Password:    "lpxiang",
		DB:          0, // use default DB
		DialTimeout: (20 * time.Second),
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(pong)

	return client
}
