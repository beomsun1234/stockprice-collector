package config_test

import (
	"fmt"
	"github/beomsun1234/stockprice-collector/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCofig(t *testing.T) {
	t.Run("config test", func(t *testing.T) {
		//given
		c := config.NewConfig()
		//when, then
		assert.Equal(t, &config.Config{}, c)
	})
}

func Test_SetCofig(t *testing.T) {
	t.Run("config setting test", func(t *testing.T) {
		//given
		c := config.NewConfig()
		//when
		c.SetConfig("test.yaml")
		// then
		fmt.Println(c.KisConfig.Key, c.KisConfig.Secret)
		fmt.Println("------------------------------------")
		assert.Equal(t, "key", c.KisConfig.Key)
		assert.Equal(t, "secret", c.KisConfig.Secret)
	})
}
