package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.nhnent.com/godo/cfo/server/log"
)

// authoirzed() if session[IsLogin] != true
// then redirect to login page with 302(found)
// or next to context
func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		isLogin := session.Get("IsLogin")

		defer log.Debug("Authorized()", session.Get("UserName"), session.Get("IsLogin"))
		if isLogin != true {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// defer log.Debug("access", session.Get("UserName"))
		c.Next()
	}
}
