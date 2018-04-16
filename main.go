package main

import (
	"code.ysitd.cloud/component/aviation/sky/pkg/bootstrap"
	"code.ysitd.cloud/component/aviation/sky/pkg/cache"
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

	if !cache.SingleNode {
		go func() {
			for {
				time.Sleep(cache.UpdateInterval)
				cache.UpdatePool(bootstrap.Logger.WithField("source", "groupcache"))
			}
		}()
		go cache.Listen(bootstrap.Logger.WithField("source", "groupcache_http"))
	}

	app.Run()
}
