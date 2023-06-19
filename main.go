package main

import (
	"github/beomsun1234/stockprice-collector/config"
	"github/beomsun1234/stockprice-collector/external/kis"
	"github/beomsun1234/stockprice-collector/infra/database"
	"github/beomsun1234/stockprice-collector/repository"
	"github/beomsun1234/stockprice-collector/scheduler"
	"github/beomsun1234/stockprice-collector/service"
	"net/http"
	"os"

	"github.com/robfig/cron"
)

var stockPriceCollectionScheduler scheduler.StockPriceCollectionSchedulerInterface

func init() {
	workingDir, _ := os.Getwd()
	c := config.NewConfig()
	c.SetConfig(workingDir + "/config/" + "properties.yaml")
	redis := database.NewRedisDB(c.RedisConfig)
	redisClinet, err := redis.Connect()
	if err != nil {
		panic(err)
	}
	s := service.NewStockPriceColletorService(kis.NewKisClientSetvice(&http.Client{}, &c.KisConfig, repository.NewKisAccessTokenRepository(redisClinet)))
	stockPriceCollectionScheduler = scheduler.NewStockPriceCollectionScheduler(s, cron.New())
}

func main() {
	stockPriceCollectionScheduler.CollectStockPricesEverySecond()
	select {}
}
