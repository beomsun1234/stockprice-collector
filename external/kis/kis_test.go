package kis_test

import (
	"github/beomsun1234/stockprice-collector/config"
	"github/beomsun1234/stockprice-collector/external/kis"
	"github/beomsun1234/stockprice-collector/external/kis/dto"
	"github/beomsun1234/stockprice-collector/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetStockPrice(t *testing.T) {
	t.Run("GetStcokPrice test", func(t *testing.T) {
		//given
		kis_config := &config.KisConfig{
			Key:    "test",
			Secret: "test",
		}
		res := &dto.KisStockPriceResponse{
			KisStockPriceResDetails: dto.KisStockPriceResponseDetails{
				Stck_Prpr: "10000",
			},
		}
		httpClient := &mocks.MockStockPriceHttpClient{}
		httpClient.MockKisStockPriceResponse = res

		kisClient := kis.NewKisClientSetvice(httpClient, kis_config)
		//when
		stock, err := kisClient.GetStockPrice("01234")
		//then
		assert.Equal(t, stock.Stock_Price, "10000")
		assert.Equal(t, err, nil)
	})
}

func Test_GetAccessToken(t *testing.T) {
	t.Run("GetAccessToken test", func(t *testing.T) {
		//given
		kis_config := &config.KisConfig{
			Key:    "test",
			Secret: "test",
		}
		res := &dto.KisAccessTokenResponse{
			AccessToken: "test_access_token",
			TokenType:   "test",
			ExpiresIn:   3,
		}
		mockAccessTokenHttpClient := &mocks.MockAccessTokenHttpClient{}
		mockAccessTokenHttpClient.MockKisTokenResponse = res
		kisClient := kis.NewKisClientSetvice(mockAccessTokenHttpClient, kis_config)
		//when
		kis_token_res, err := kisClient.GetAccesstoken()
		//then
		assert.Equal(t, kis_token_res.AccessToken, "test_access_token")
		assert.Equal(t, kis_token_res.TokenType, "test")
		assert.Equal(t, kis_token_res.ExpiresIn, 3)
		assert.Equal(t, err, nil)
	})
}
