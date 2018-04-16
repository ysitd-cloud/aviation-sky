package cache

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/golang/groupcache"
	"github.com/sirupsen/logrus"
)

const lookupTimeout = 5 * time.Second
const UpdateInterval = 1 * time.Minute

func UpdatePool(logger logrus.FieldLogger) {
	pool := Picker.(*groupcache.HTTPPool)
	dns := os.Getenv("SERVICE_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), lookupTimeout)
	defer cancel()
	results, err := net.DefaultResolver.LookupIPAddr(ctx, dns)
	if err != nil {
		logger.Error(err)
		return
	}
	peers := make([]string, len(results))
	for _, result := range results {
		peers = append(peers, fmt.Sprintf("http://%s:50005", result.IP.String()))
	}
	pool.Set(peers...)
}
