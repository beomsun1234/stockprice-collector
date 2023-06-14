package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	KisConfig KisConfig `yaml:"kis"`
}

type KisConfig struct {
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetConfig(file_name string) {
	new_config := ConvertYamlToConfig(file_name)
	c.KisConfig = new_config.KisConfig
}

func ConvertYamlToConfig(file_name string) *Config {
	properties, err := ioutil.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	config := &Config{}
	err = yaml.Unmarshal(properties, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
