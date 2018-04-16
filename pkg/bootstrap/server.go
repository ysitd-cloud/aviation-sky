package bootstrap

import (
	"net/http"

	gracehttp "code.ysitd.cloud/component/aviation/sky/pkg/grace/http"
)

var app *gracehttp.App

func initApp() {
	logger := initLogger()
	service := initService(logger)
	app = gracehttp.New([]*http.Server{
		service.CreateServer(),
	})
	app.Logger = logger.WithField("source", "gracehttp")
}
