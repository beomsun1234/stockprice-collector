package kis_test

import (
	"github/beomsun1234/stockprice-collector/config"
	"github/beomsun1234/stockprice-collector/external/kis"
	"github/beomsun1234/stockprice-collector/external/kis/dto"
	"github/beomsun1234/stockprice-collector/mocks"
	"strconv"
	"testing"
	"time"

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
		mockRepo := mocks.NewMockKisAccessTokenRepository()
		kisClient := kis.NewKisClientSetvice(httpClient, kis_config, mockRepo)
		//when
		stock, err := kisClient.GetStockPrice("01234", "test")
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
		mockRepo := mocks.NewMockKisAccessTokenRepository()
		kisClient := kis.NewKisClientSetvice(mockAccessTokenHttpClient, kis_config, mockRepo)
		//when
		kis_token_res, err := kisClient.GetAccesstoken()
		//then
		assert.Equal(t, kis_token_res.AccessToken, "test_access_token")
		assert.Equal(t, kis_token_res.ExpiresIn, "3")
		assert.Equal(t, err, nil)
	})
}
func Test_tiem(t *testing.T) {
	t.Run("GetAccessToken test", func(t *testing.T) {
		tokenIssedAt, _ := time.Parse("2006-01-02 15:04:05", "2023-06-17 23:26:00")
		expiresIn, _ := strconv.ParseInt("86400", 10, 64)
		expiredAt := tokenIssedAt.Add(time.Second * time.Duration(expiresIn))
		now := time.Now()
		assert.Equal(t, expiredAt.After(now), true)
	})
}
