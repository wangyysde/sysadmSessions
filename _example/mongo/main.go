package main

import (
	"github.com/wangyysde/sysadmSessions"
	"github.com/wangyysde/sysadmSessions/mongo"
	"github.com/wangyysde/sysadmServer"
	"github.com/globalsign/mgo"
)

func main() {
	r := sysadmServer.Default()
	session, err := mgo.Dial("localhost:27017/test")
	if err != nil {
		// handle err
	}

	c := session.DB("").C("sysadmSessions")
	store := mongo.NewStore(c, 3600, true, []byte("secret"))
	r.Use(sysadmSessions.Sessions("mysession", store))

	r.GET("/incr", func(c *sysadmServer.Context) {
		session := sysadmSessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, sysadmServer.H{"count": count})
	})
	r.Run(":8000")
}
