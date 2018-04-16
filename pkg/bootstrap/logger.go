package bootstrap

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func initLogger() *logrus.Logger {
	Logger = logrus.New()

	if os.Getenv("VERBOSE") != "" {
		Logger.SetLevel(logrus.DebugLevel)
	}

	return Logger
}
