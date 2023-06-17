package repository_test

import (
	"encoding/json"
	"github/beomsun1234/stockprice-collector/domain"
	"github/beomsun1234/stockprice-collector/mocks"
	"github/beomsun1234/stockprice-collector/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAccessToken(t *testing.T) {
	t.Run("Test GetAccessToken", func(t *testing.T) {
		//given
		db := mocks.NewMockRedisDatabase()
		db.Connect()
		repo := repository.NewKisAccessTokenRepository(db)
		token := domain.NewToken().BuildAccessToken("test").BuildExpiresIn("10").BuildIssuedAt("2023-11-16 11:00:00")
		byte_token, _ := json.Marshal(token)
		db.RedisMock.ExpectHGet("token", "token").SetVal(string(byte_token))
		//when
		res_token, err := repo.GetKisAccessToken()
		//then
		assert.Equal(t, token.AccessToken, res_token.AccessToken)
		assert.NoError(t, err)
	})
}

func Test_InsertAccessToken(t *testing.T) {
	t.Run("Test InsertAccessToken", func(t *testing.T) {
		//given
		db := mocks.NewMockRedisDatabase()
		db.Connect()
		repo := repository.NewKisAccessTokenRepository(db)

		tt := domain.NewToken().BuildAccessToken("test").BuildExpiresIn("10").BuildIssuedAt("2023-11-16 11:00:00")

		data, _ := json.Marshal(tt)

		db.RedisMock.ExpectHSet("token", "token", data).SetVal(1)
		//when
		err := repo.InsertKisAccessToken(tt)
		//then
		assert.Nil(t, err)
	})
}

func Test_DeleteAccessToken(t *testing.T) {
	t.Run("Test DeleteAccessToken", func(t *testing.T) {
		//given
		db := mocks.NewMockRedisDatabase()
		db.Connect()
		repo := repository.NewKisAccessTokenRepository(db)

		domain.NewToken().BuildAccessToken("test").BuildExpiresIn("10").BuildIssuedAt("2023-11-16 11:00:00")

		db.RedisMock.ExpectHDel("token", "token")
		//when,then
		repo.DeleteKisAccessToken()

	})
}
