package config

import (
	"github.com/spf13/viper"
	"log"
)

func Init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

// TODO: Add functions to get each value/section of the config file.
