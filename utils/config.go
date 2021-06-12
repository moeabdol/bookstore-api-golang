package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ConfigStruct type
type ConfigStruct struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	Port        string `mapstructure:"PORT"`
	DBDialect   string `mapstructure:"DB_DIALECT"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBSslmode   string `mapstructure:"DB_SSLMODE"`
	DBName      string `mapstructure:"DB_NAME"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
}

// Config global variable
var Config *ConfigStruct

// LoadConfig function to read .env configuration file
func LoadConfig() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found!")
		} else {
			log.Fatal("Something went wrong while loading config file!")
		}
	}

	viper.AutomaticEnv() // Override config file with environment variables

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatal("Unable to read config file")
	}
}
