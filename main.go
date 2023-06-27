package main

import (
	"github/beomsun1234/stockprice-collector/config"
	"github/beomsun1234/stockprice-collector/external/kis"
	"github/beomsun1234/stockprice-collector/infra/database"
	"github/beomsun1234/stockprice-collector/infra/messagequeue"
	"github/beomsun1234/stockprice-collector/repository"
	"github/beomsun1234/stockprice-collector/scheduler"
	"github/beomsun1234/stockprice-collector/service"
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/robfig/cron"
)

var stockPriceCollectionScheduler scheduler.StockPriceCollectionSchedulerInterface

func init() {
	c := config.NewConfig()
	c.SetConfig("properties.yaml")
	redis := database.NewRedisDB(c.RedisConfig)
	redisClinet, err := redis.Connect()
	if err != nil {
		panic(err)
	}
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaClient, err := sarama.NewClient(c.KafkaConfig.Addrs, kafkaConfig)
	if err != nil {
		panic(err)
	}
	kafka := messagequeue.NewKafka(&kafkaClient)
	s := service.NewStockPriceColletorService(kis.NewKisClientSetvice(&http.Client{}, &c.KisConfig, repository.NewKisAccessTokenRepository(redisClinet)))
	stockPriceCollectionScheduler = scheduler.NewStockPriceCollectionScheduler(s, cron.New(), kafka)
}

func main() {
	stockPriceCollectionScheduler.CollectStockPricesEverySecond()
	select {}
}
