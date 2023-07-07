package common

import (
	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func InitRedis() *redis.Client {
	return rdb
}

// *redis
func GetRedis() *redis.Client {
	return rdb
}
