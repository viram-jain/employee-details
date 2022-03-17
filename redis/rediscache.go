package redis

import (
	"encoding/json"
	"fmt"
	"os"

	"time"

	"github.com/go-redis/redis"
	"github.com/viram-jain/employee-details/model"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value model.Employee) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	client.Set(key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *model.Employee {
	client := cache.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	post := model.Employee{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &post
}
