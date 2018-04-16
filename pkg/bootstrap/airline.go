package bootstrap

import (
	"code.ysitd.cloud/component/aviation/sky/pkg/modals/airline"
	"github.com/dgrijalva/lfu-go"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	"os"
)

func initAirline(logger logrus.FieldLogger) airline.Store {
	s3Endpoint := os.Getenv("S3_ENDPOINT")
	if s3Endpoint == "" {
		s3Endpoint = "s3.amazonaws.com"
	}

	client, _ := minio.New(s3Endpoint, os.Getenv("S3_ACCESS_KEY_ID"), os.Getenv("S3_SECRET_ACCESS_KEY"), true)

	cache := lfu.New()
	cache.UpperBound = 100
	cache.LowerBound = 60

	blob := airline.NewBlobStore(client, os.Getenv("S3_BUCKET"), "airline")
	blob.Logger = logger.WithField("source", "airline_blob")

	return &airline.PluginStore{
		BlobStore: blob,
		Cache:     cache,
		Logger:    logger.WithField("source", "airline_store"),
	}
}
