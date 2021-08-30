package routers

import (
	"github.com/gin-gonic/gin"
	"go_gin_study/server"
	"go_gin_study/service"
	"net/http"
)

func Users(c *gin.Context) {
	userService := service.User{}
	users, err := userService.GetAll()
	server.Debug(users, err)
	if err != nil {
		return
	}

	c.String(http.StatusOK, "Users success")
}

func Signup(c *gin.Context)  {

}