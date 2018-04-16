package bootstrap

import (
	"code.ysitd.cloud/component/aviation/sky/pkg/server"
	"github.com/sirupsen/logrus"
)

func initService(logger logrus.FieldLogger) *server.Service {
	return &server.Service{
		Logger:   logger.WithField("source", "service"),
		Hostname: initFlyerStore(logger),
		Airline:  initAirline(logger),
	}
}
