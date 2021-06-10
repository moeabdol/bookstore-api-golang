package utils

import (
	"log"

	"github.com/spf13/viper"
)

// ReadConfig function to read .env file
func ReadConfig() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found!")
		} else {
			log.Fatal("Something went wrong while loading config file!")
		}
	}
}
