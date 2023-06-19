package database

import (
	"errors"
	"github/beomsun1234/stockprice-collector/config"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Redis       *redis.Client
	RedisConfig config.RedisConfig
}

func NewRedisDB(redisConfig config.RedisConfig) *Redis {
	return &Redis{
		RedisConfig: redisConfig,
	}
}

func (r *Redis) Connect() (*redis.Client, error) {
	r.Redis = redis.NewClient(&redis.Options{
		Addr:     r.RedisConfig.Addr,
		Password: r.RedisConfig.Password, // no password set
		// use default DB
	})
	if r.Redis != nil {
		return nil, errors.New("redis connection error")
	}
	return r.Redis, nil
}
