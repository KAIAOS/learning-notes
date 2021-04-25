package dao

import (
	"github.com/go-redis/redis"
)

var Redisdb *redis.Client

func InitRedis()(err error){
	Redisdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	_, err = Redisdb.Ping().Result()
	return
}