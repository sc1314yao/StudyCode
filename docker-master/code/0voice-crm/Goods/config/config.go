package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Goods struct {
		Name     string
		Category string
	}
}

var conf *Config

func init() {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile("config.yaml")
	v.ReadInConfig()
	conf = &Config{}
	err := v.Unmarshal(conf)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConf() *Config {
	return conf
}
