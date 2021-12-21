package mongo

import (
	"github.com/wangyysde/sysadmSessions"
	"github.com/globalsign/mgo"
	"github.com/kidstuff/mongostore"
)

type Store interface {
	sysadmSessions.Store
}

func NewStore(c *mgo.Collection, maxAge int, ensureTTL bool, keyPairs ...[]byte) Store {
	return &store{mongostore.NewMongoStore(c, maxAge, ensureTTL, keyPairs...)}
}

type store struct {
	*mongostore.MongoStore
}

func (c *store) Options(options sysadmSessions.Options) {
	c.MongoStore.Options = options.ToGorillaOptions()
}
