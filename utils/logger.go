package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Log is the logger global variable
var Log *log.Logger

// InitializeLogger function
func InitializeLogger() {
	Log = log.New()
	Log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
		PadLevelText:  true,
	})

	environment := viper.Get("ENVIRONMENT").(string)
	if environment == "development" {
		Log.SetLevel(log.DebugLevel)
		Log.Info("Server running in development mode")
	} else {
		Log.SetLevel(log.InfoLevel)
		Log.Info("Server running in production mode")
	}

	Log.Info("Finished initializeing logger")
}
