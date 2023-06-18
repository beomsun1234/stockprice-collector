package repository

import (
	"context"
	"encoding/json"
	"github/beomsun1234/stockprice-collector/domain"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type KisAccessTokenRepositoryInterface interface {
	GetKisAccessToken() (*domain.Token, error)
	DeleteKisAccessToken()
	InsertKisAccessToken(token *domain.Token) error
}

type KisAccessTokenRepository struct {
	Redis *redis.Client
}

func NewKisAccessTokenRepository(redis *redis.Client) KisAccessTokenRepositoryInterface {
	return &KisAccessTokenRepository{
		Redis: redis,
	}
}

func (k *KisAccessTokenRepository) GetKisAccessToken() (*domain.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	data, err := k.Redis.HGet(ctx, "token", "token").Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	token := domain.NewToken()
	json.Unmarshal([]byte(data), token)
	return token, nil
}

func (k *KisAccessTokenRepository) DeleteKisAccessToken() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	k.Redis.Del(ctx, "token")
	defer cancel()
}

func (k *KisAccessTokenRepository) InsertKisAccessToken(token *domain.Token) error {
	data, _ := json.Marshal(token)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	err := k.Redis.HSet(ctx, "token", "token", data).Err()
	defer cancel()
	if err != nil {
		log.Println("redis insert err", err)
		return err
	}
	return nil
}
