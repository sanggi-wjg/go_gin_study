package home

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin_study/service"
	"net/http"
)

func Ping(c *gin.Context) {
	userService := service.User{}
	users, err := userService.GetAll()
	fmt.Println(users)
	fmt.Println(err)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
