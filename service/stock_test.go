package service_test

import (
	"github/beomsun1234/stockprice-collector/config"
	"github/beomsun1234/stockprice-collector/external/kis"
	"github/beomsun1234/stockprice-collector/external/kis/dto"
	"github/beomsun1234/stockprice-collector/mocks"
	"github/beomsun1234/stockprice-collector/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CollectStockPrices(t *testing.T) {
	t.Run("CollectStockPrices Test", func(t *testing.T) {
		//given
		c := config.NewConfig()
		c.KisConfig = config.KisConfig{
			Key:    "key",
			Secret: "secret",
		}
		httpClient := &mocks.MockStockPriceHttpClient{}
		kisClient := kis.NewKisClientSetvice(httpClient, &c.KisConfig)
		stockPriceCollectorService := service.NewStockPriceColletorService(kisClient)
		res := &dto.KisStockPriceResponse{
			KisStockPriceResDetails: dto.KisStockPriceResponseDetails{
				Stck_Prpr: "10000",
			},
		}
		httpClient.MockKisStockPriceResponse = res
		//when
		stocks := stockPriceCollectorService.CollectStockPrices()
		//then
		response_size := len(stocks)
		request_stock_size := len(service.GetStockCodes())
		assert.Equal(t, request_stock_size, response_size)
		assert.Equal(t, "10000", stocks[0].Stock_Price)
	})

}
