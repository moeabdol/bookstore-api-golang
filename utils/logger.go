package utils

import log "github.com/sirupsen/logrus"

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

	if Config.Environment == "development" {
		Log.SetLevel(log.DebugLevel)
		Log.Info("Server running in development mode")
	} else {
		Log.SetLevel(log.InfoLevel)
		Log.Info("Server running in production mode")
	}

	Log.Info("Finished initializeing logger")
}
