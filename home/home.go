package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	//users, _ := models.GetUserList()
	//fmt.Println(users)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
