package database

import (
	"github/beomsun1234/stockprice-collector/config"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Redis       *redis.Client
	RedisConfig config.RedisConfig
}

func NewRedisDB(redisConfig config.RedisConfig) Database {
	return &Redis{
		RedisConfig: redisConfig,
	}
}

func (r *Redis) Client() *redis.Client {
	return r.Redis
}

func (r *Redis) Connect() {
	r.Redis = redis.NewClient(&redis.Options{
		Addr:     r.RedisConfig.Addr,
		Password: r.RedisConfig.Password, // no password set
		// use default DB
	})
}
