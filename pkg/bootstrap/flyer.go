package bootstrap

import (
	"os"

	"code.ysitd.cloud/common/go/db"
	"code.ysitd.cloud/component/aviation/sky/pkg/modals/flyer"
	"github.com/sirupsen/logrus"
)

func initFlyerStore(logger logrus.FieldLogger) *flyer.Store {
	dbPool := db.NewPool("postgres", os.Getenv("DB_URL"))
	return flyer.NewStore(dbPool, logger.WithField("source", "flyer_store"))
}
