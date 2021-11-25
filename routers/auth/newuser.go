package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.nhnent.com/godo/cfo/server/log"
	"github.nhnent.com/godo/cfo/service/user_service"
)

/*
@Success 200 / newuser template page
*/
func GetNewUser(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/newuser", gin.H{
		"title": "Join",
	})
}

/*
@Form
@Success 200 / return success
@Failure 400 / return error
*/
type newUserJoinForm struct {
	UserName  string `form:"username"  binding:"required"`
	Password  string `form:"password"  binding:"required"`
	UserLevel uint   `form:"userlevel"  binding:"required"`
}

func PostNewUser(c *gin.Context) {
	var user newUserJoinForm
	if err := c.Bind(&user); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userService := user_service.User{UserName: user.UserName, Password: user.Password, UserLevel: user.UserLevel}
	id, err := userService.CreateUser()
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"id":     id,
	})
}
