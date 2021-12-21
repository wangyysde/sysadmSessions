package memcached

import (
	"github.com/bradfitz/gomemcache/memcache"
	gsm "github.com/bradleypeabody/gorilla-sessions-memcache"
	"github.com/wangyysde/sysadmSessions"
)

type Store interface {
	sysadmSessions.Store
}

// client: memcache client (github.com/bradfitz/gomemcache/memcache)
// keyPrefix: prefix for the keys we store.
func NewStore(
	client *memcache.Client, keyPrefix string, keyPairs ...[]byte,
) Store {
	memcacherClient := gsm.NewGoMemcacher(client)
	return NewMemcacheStore(memcacherClient, keyPrefix, keyPairs...)
}

// client: memcache client which implements the gsm.Memcacher interface
// keyPrefix: prefix for the keys we store.
func NewMemcacheStore(
	client gsm.Memcacher, keyPrefix string, keyPairs ...[]byte,
) Store {
	return &store{gsm.NewMemcacherStore(client, keyPrefix, keyPairs...)}
}

type store struct {
	*gsm.MemcacheStore
}

func (c *store) Options(options sysadmSessions.Options) {
	c.MemcacheStore.Options = options.ToGorillaOptions()
}
