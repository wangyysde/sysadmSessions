package mongo

import (
	"testing"

	"github.com/wangyysde/sysadmSessions"
	"github.com/wangyysde/sysadmSessions/tester"
	"github.com/globalsign/mgo"
)

const mongoTestServer = "localhost:27017"

var newStore = func(_ *testing.T) sysadmSessions.Store {
	session, err := mgo.Dial(mongoTestServer)
	if err != nil {
		panic(err)
	}

	c := session.DB("test").C("sysadmSessions")
	return NewStore(c, 3600, true, []byte("secret"))
}

func TestMongo_SessionGetSet(t *testing.T) {
	tester.GetSet(t, newStore)
}

func TestMongo_SessionDeleteKey(t *testing.T) {
	tester.DeleteKey(t, newStore)
}

func TestMongo_SessionFlashes(t *testing.T) {
	tester.Flashes(t, newStore)
}

func TestMongo_SessionClear(t *testing.T) {
	tester.Clear(t, newStore)
}

func TestMongo_SessionOptions(t *testing.T) {
	tester.Options(t, newStore)
}

func TestMongo_SessionMany(t *testing.T) {
	tester.Many(t, newStore)
}
