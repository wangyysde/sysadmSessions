// +build go1.11

package tester

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/wangyysde/sysadmSessions"
	"github.com/wangyysde/sysadmServer"
)

func testOptionSameSitego(t *testing.T, r *sysadmServer.Engine) {
	r.GET("/sameSite", func(c *sysadmServer.Context) {
		session := sysadmSessions.Default(c)
		session.Set("key", ok)
		session.Options(sysadmSessions.Options{
			SameSite: http.SameSiteStrictMode,
		})
		_ = session.Save()
		c.String(200, ok)
	})

	res3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("GET", "/sameSite", nil)
	r.ServeHTTP(res3, req3)

	s := strings.Split(res3.Header().Get("Set-Cookie"), ";")
	if s[1] != " SameSite=Strict" {
		t.Error("Error writing samesite with options:", s[1])
	}
}
