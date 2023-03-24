package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Zaim Zaim `yaml:"zaim"`
}

type Zaim struct {
	ConsumerKey string `yaml:"consumerKey"`
	ConsumerSecret string `yaml:"consumerSecret"`
	AccessToken string `yaml:"accessToken"`
	AccessSecret string `yaml:"accessSecret"`
}

var Conf = NewConfig()

func NewConfig() Config {
	viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return config
}
