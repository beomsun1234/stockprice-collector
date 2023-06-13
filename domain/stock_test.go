package domain_test

import (
	"github/beomsun1234/stockprice-collector/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BuildTest(t *testing.T) {
	t.Run("Stock Build Test", func(t *testing.T) {
		//given
		s := domain.NewStock()
		//then
		s.BuildStockCode("1234").BuildStockPrice("123123").BuildStockHighestPrice("1").BuildStockLowestPrice("1")
		s.BuildStockVolume("1").BuildStockPrdyVrssSign("1")
		//when
		assert.Equal(t, s.Stock_Code, "1234")
		assert.Equal(t, s.Stock_Price, "123123")
		assert.Equal(t, s.Stock_Lowest_Price, "1")
		assert.Equal(t, s.Stock_Highest_Price, "1")
		assert.Equal(t, s.Stock_Volume, "1")
		assert.Equal(t, s.Stock_Prdy_Vrss_Sign, "1")
	})

}
