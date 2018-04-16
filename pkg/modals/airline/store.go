package airline

import (
	"context"
	"io/ioutil"
	"os"

	"code.ysitd.cloud/component/aviation/runway"
	"code.ysitd.cloud/component/aviation/runway/validate"
	"github.com/dgrijalva/lfu-go"
	"github.com/sirupsen/logrus"
)

type PluginStore struct {
	BlobStore *BlobStore
	Cache     *lfu.Cache
	Logger    logrus.FieldLogger
}

func (ps *PluginStore) GetRevision(ctx context.Context, revision string) (airline runway.Airline, err error) {
	val := ps.Cache.Get(revision)
	if val != nil {
		return val.(runway.Airline), nil
	}

	airline, err = ps.installPlugin(ctx, revision)
	if err != nil {
		return
	}
	ps.Cache.Set(revision, airline)
	return
}

func (ps *PluginStore) installPlugin(ctx context.Context, revision string) (airline runway.Airline, err error) {
	bytes, err := ps.BlobStore.Load(ctx, revision)
	if err != nil {
		return
	}

	file, err := ioutil.TempFile("", "airline"+revision)
	if err != nil {
		return
	}

	if err != nil {
		return
	}

	defer os.RemoveAll(file.Name())

	if _, err := file.Write(bytes); err != nil {
		return nil, err
	}

	if err := file.Close(); err != nil {
		return nil, err
	}

	airline, err = validate.ValidateAirline(file.Name())
	if err != nil {
		return
	}

	airline.Initial(ctx)

	ps.Logger.Debugf("Install plugin with revision %s\n", revision)
	return
}
