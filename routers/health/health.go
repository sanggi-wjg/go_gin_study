package health

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	session := sessions.Default(c)
	fmt.Println(session.Get("UserName"), session.ID())
	// time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
