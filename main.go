package main

import (
	"code.ysitd.cloud/component/aviation/sky/pkg/bootstrap"
	"time"
)

const restartTick = time.Minute * 60

func main() {
	app := bootstrap.GetApp()
	go func() {
		for {
			time.Sleep(restartTick)
			app.Restart()
		}
	}()
	app.Run()
}
