package domain_test

import (
	"github/beomsun1234/stockprice-collector/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_TokenTest(t *testing.T) {
	t.Run("token test", func(t *testing.T) {
		now := time.Now().Format("2006-01-02 15:04:05")
		token := domain.NewToken().BuildAccessToken("1234").BuildExpiresIn("86400").BuildIssuedAt(now)

		assert.Equal(t, false, token.IsTokenExpired())

	})
}
