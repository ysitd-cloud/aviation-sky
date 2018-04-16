package airline

import (
	"bytes"
	"context"
	"io"
	"time"

	"github.com/golang/groupcache"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
)

const timeout = 10 * time.Second
const cacheSize int64 = 32 << 20

type BlobStore struct {
	Cache  *groupcache.Group
	Client *minio.Client
	Bucket string
	Logger logrus.FieldLogger
}

func NewBlobStore(client *minio.Client, bucket, groupName string) *BlobStore {
	store := &BlobStore{
		Client: client,
		Bucket: bucket,
	}

	group := groupcache.GetGroup(groupName)
	if group == nil {
		group = groupcache.NewGroup(groupName, cacheSize, store)
	}

	store.Cache = group

	return store
}

func (s *BlobStore) Get(_ groupcache.Context, key string, dest groupcache.Sink) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	obj, err := s.Client.GetObjectWithContext(ctx, s.Bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return
	}

	defer obj.Close()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, obj)
	if err != nil {
		return
	}

	s.Logger.Debugf("Get file for revision %s from remote", key)

	return dest.SetBytes(buffer.Bytes())
}

func (s *BlobStore) Load(ctx context.Context, revision string) (bytes []byte, err error) {
	err = s.Cache.Get(ctx, revision, groupcache.AllocatingByteSliceSink(&bytes))
	return
}
