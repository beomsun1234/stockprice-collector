package database

import "github.com/go-redis/redis/v8"

type Database interface {
	Connect()
	Client() *redis.Client
}
