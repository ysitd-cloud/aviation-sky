package cache

import (
	"github.com/golang/groupcache"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Listen(logger logrus.FieldLogger) {
	pool := Picker.(*groupcache.HTTPPool)
	logger.Error(http.ListenAndServe(":50005", pool))
}
