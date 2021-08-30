package routers

import (
	"github.com/gin-gonic/gin"
	"go_gin_study/server"
	"go_gin_study/service"
	"net/http"
)

func Ping(c *gin.Context) {
	userService := service.User{}
	users, err := userService.GetAll()
	server.Debug(users, err)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
