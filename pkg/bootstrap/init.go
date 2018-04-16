package bootstrap

import "code.ysitd.cloud/component/aviation/sky/pkg/grace/http"

func init() {
	initApp()
}

func GetApp() *http.App {
	return app
}
