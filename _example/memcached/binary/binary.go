package main

import (
	"github.com/wangyysde/sysadmSessions"
	"github.com/wangyysde/sysadmSessions/memcached"
	"github.com/wangyysde/sysadmServer"
	"github.com/memcachier/mc"
)

func main() {
	r := sysadmServer.Default()
	client := mc.NewMC("localhost:11211", "username", "password")
	store := memcached.NewMemcacheStore(client, "", []byte("secret"))
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
