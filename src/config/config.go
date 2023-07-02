package config

import (
	"github.com/spf13/viper"
)

func InitConfig() (*viper.Viper, error) {
	conf := viper.New()
	conf.SetConfigName("env.yaml")
	conf.SetConfigType("yaml")
	conf.AddConfigPath("./src/config")
	conf.AutomaticEnv()
	err := conf.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return conf, nil
}
