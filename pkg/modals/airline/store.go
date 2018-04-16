package airline

import (
	"context"
	"io/ioutil"
	"os"

	"code.ysitd.cloud/component/aviation/runway"
	"code.ysitd.cloud/component/aviation/runway/validate"
	"github.com/dgrijalva/lfu-go"
)

type PluginStore struct {
	BlobStore *BlobStore
	Cache     *lfu.Cache
}

func (ps *PluginStore) GetRevision(ctx context.Context, revision string) (runway.Airline, error) {
	val := ps.Cache.Get(revision)
	if val != nil {
		return val.(runway.Airline), nil
	}

	return ps.installPlugin(ctx, revision)
}

func (ps *PluginStore) installPlugin(ctx context.Context, revision string) (airline runway.Airline, err error) {
	bytes, err := ps.BlobStore.Load(ctx, revision)
	if err != nil {
		return
	}

	file, err := ioutil.TempFile("airline", revision)
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
	return
}
