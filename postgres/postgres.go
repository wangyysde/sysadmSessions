package postgres

import (
	"database/sql"
	"github.com/antonlindstrom/pgstore"
	"github.com/wangyysde/sysadmSessions"
)

type Store interface {
	sysadmSessions.Store
}

type store struct {
	*pgstore.PGStore
}

var _ Store = new(store)

func NewStore(db *sql.DB, keyPairs ...[]byte) (Store, error) {
	p, err := pgstore.NewPGStoreFromPool(db, keyPairs...)
	if err != nil {
		return nil, err
	}

	return &store{p}, nil
}

func (s *store) Options(options sysadmSessions.Options) {
	s.PGStore.Options = options.ToGorillaOptions()
}
