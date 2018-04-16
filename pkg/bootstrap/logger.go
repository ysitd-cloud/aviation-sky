package bootstrap

import (
	"github.com/sirupsen/logrus"
	"os"
)

func initLogger() *logrus.Logger {
	logger := logrus.New()

	if os.Getenv("VERBOSE") != "" {
		logger.SetLevel(logrus.DebugLevel)
	}

	return logger
}
