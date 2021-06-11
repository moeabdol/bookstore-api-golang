package utils

import (
	"log"

	"github.com/spf13/viper"
)

// ConfigStruct type
type ConfigStruct struct {
	Environment string
	Port        string
	DbDialect   string
	DbHost      string
	DbPort      string
	DbSslMode   string
	DbName      string
	DbUser      string
	DbPassword  string
}

// Config global variable
var Config ConfigStruct

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

	Config.Environment = viper.Get("Environment").(string)
	Config.Port = viper.Get("PORT").(string)
	Config.DbDialect = viper.Get("DB_DIALECT").(string)
	Config.DbHost = viper.Get("DB_HOST").(string)
	Config.DbPort = viper.Get("DB_PORT").(string)
	Config.DbSslMode = viper.Get("DB_SSLMODE").(string)
	Config.DbName = viper.Get("DB_NAME").(string)
	Config.DbUser = viper.Get("DB_USER").(string)
	Config.DbPassword = viper.Get("DB_PASSWORD").(string)
}
