package bootstrap

import (
	"code.ysitd.cloud/component/aviation/sky/pkg/server"
	"github.com/sirupsen/logrus"
)

var service *server.Service

func initService(logger logrus.FieldLogger) *server.Service {
	service = &server.Service{
		Logger:   logger.WithField("source", "service"),
		Hostname: initFlyerStore(logger),
		Airline:  initAirline(logger),
	}
	return service
}
