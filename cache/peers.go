package cache

import "cache_go/cache/cachepb"

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(request *cachepb.Request, response *cachepb.Response) error
}
