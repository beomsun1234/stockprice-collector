package scheduler_test

import (
	"context"
	"github/beomsun1234/stockprice-collector/config"
	"github/beomsun1234/stockprice-collector/external/kis"
	"github/beomsun1234/stockprice-collector/external/kis/dto"
	"github/beomsun1234/stockprice-collector/mocks"
	"github/beomsun1234/stockprice-collector/scheduler"
	"github/beomsun1234/stockprice-collector/service"
	"testing"
	"time"

	"github.com/robfig/cron"
)

func Test_CollectStockPricesEverySecond(t *testing.T) {
	t.Run("scheduler test", func(t *testing.T) {
		//given
		c := config.NewConfig()
		c.KisConfig = config.KisConfig{
			Key:    "key",
			Secret: "secret",
		}
		httpClient := &mocks.MockStockPriceHttpClient{}
		mockRepo := mocks.NewMockKisAccessTokenRepository()
		kisClient := kis.NewKisClientSetvice(httpClient, &c.KisConfig, mockRepo)
		kafkaMock := mocks.NewMockMessagequeue()
		stockPriceCollectorService := service.NewStockPriceColletorService(kisClient)
		res := &dto.KisStockPriceResponse{
			KisStockPriceResDetails: dto.KisStockPriceResponseDetails{
				Stck_Prpr: "10000",
			},
		}
		httpClient.MockKisStockPriceResponse = res
		mock_scheduler := &mocks.MockScheduler{}
		s := scheduler.NewStockPriceCollectionScheduler(stockPriceCollectorService, mock_scheduler, kafkaMock)
		//when, then
		s.CollectStockPricesEverySecond()
	})
}

func Test_CollectStockPricesEverySecondReal(t *testing.T) {
	t.Run("scheduler test", func(t *testing.T) {
		//given
		c := config.NewConfig()
		c.KisConfig = config.KisConfig{
			Key:    "key",
			Secret: "secret",
		}
		httpClient := &mocks.MockStockPriceHttpClient{}
		mockRepo := mocks.NewMockKisAccessTokenRepository()
		kafkaMock := mocks.NewMockMessagequeue()
		kisClient := kis.NewKisClientSetvice(httpClient, &c.KisConfig, mockRepo)
		stockPriceCollectorService := service.NewStockPriceColletorService(kisClient)
		res := &dto.KisStockPriceResponse{
			KisStockPriceResDetails: dto.KisStockPriceResponseDetails{
				Stck_Prpr: "10000",
			},
		}
		httpClient.MockKisStockPriceResponse = res
		scheduler_di := cron.New()
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		s := scheduler.NewStockPriceCollectionScheduler(stockPriceCollectorService, scheduler_di, kafkaMock)
		//when, then
		s.CollectStockPricesEverySecond()
		<-ctx.Done()

		defer cancel()
	})
}
