package bootstrap

import (
	"code.ysitd.cloud/component/aviation/sky/pkg/grace/http"
	"code.ysitd.cloud/component/aviation/sky/pkg/server"
)

func init() {
	initApp()
}

func GetApp() *http.App {
	return app
}

func GetService() *server.Service {
	return service
}
