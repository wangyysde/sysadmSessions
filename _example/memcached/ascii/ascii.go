package main

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/wangyysde/sysadmSessions"
	"github.com/wangyysde/sysadmSessions/memcached"
	"github.com/wangyysde/sysadmServer"
)

func main() {
	r := sysadmServer.Default()
	store := memcached.NewStore(memcache.New("localhost:11211"), "", []byte("secret"))
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
