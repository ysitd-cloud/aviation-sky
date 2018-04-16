package cache

import (
	"github.com/golang/groupcache"
	"os"
)

var Picker groupcache.PeerPicker
var SingleNode = os.Getenv("SINGLE_NODE") == ""

func init() {
	if SingleNode {
		Picker = &groupcache.NoPeers{}
	} else {
		Picker = groupcache.NewHTTPPoolOpts(os.Getenv("POD_IP"), &groupcache.HTTPPoolOptions{
			Replicas: 2,
		})
	}
}
