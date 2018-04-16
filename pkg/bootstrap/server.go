package bootstrap

import (
	"net/http"
	"os"

	gracehttp "code.ysitd.cloud/component/aviation/sky/pkg/grace/http"
)

var app *gracehttp.App

func initApp() {
	logger := initLogger()
	service := initService(logger)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app = gracehttp.New([]*http.Server{
		service.CreateServer(":" + port),
	})
	app.Logger = logger.WithField("source", "gracehttp")
}
