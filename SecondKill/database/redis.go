package database

import (
	"SecondKill/config"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

var redisInstance *redis.Database

func GetRedisInstance() *redis.Database {
	if redisInstance == nil {
		redisInstance = initRedis()
	}
	return redisInstance
}

func initRedis() *redis.Database {
	var dataBase *redis.Database

	rd := config.RedisSetting
	dataBase = redis.New(redis.Config{
		Network:     rd.NetWork,
		Addr:        rd.Host,
		Password:    rd.Password,
		Database:    "",
		MaxActive:   rd.MaxActive,
		Timeout: redis.DefaultRedisTimeout,
		Prefix:      rd.Prefix,
	})
	return dataBase
}
