package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.nhnent.com/godo/cfo/server/config"
	"github.nhnent.com/godo/cfo/server/log"
	"github.nhnent.com/godo/cfo/server/util"
	"github.nhnent.com/godo/cfo/service/user_service"
)

/*
@Success 200 / login template page
*/
func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/login", gin.H{
		"title": "Login",
	})
}

type userLoginForm struct {
	UserName string `form:"username"  binding:"required"`
	Password string `form:"password"  binding:"required"`
}

/*
@JSON
@Success 200 / return json token
@Failure 400 / return json error
*/
func PostLogin(c *gin.Context) {
	var user userLoginForm
	if err := c.Bind(&user); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userSerivce := user_service.User{UserName: user.UserName, Password: user.Password}
	res, err := userSerivce.CheckUserByUserNameAndPassword()
	if !res {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	token, err := util.GenerateToken(user.UserName, user.Password)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to generate token",
		})
		return
	}
	createSession(c, user, token)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}

// create session
func createSession(c *gin.Context, user userLoginForm, token string) {
	defer log.Debug("create session", user)
	session := sessions.Default(c)
	session.Options(sessions.Options{Path: config.AppConfig.SessPath, MaxAge: config.AppConfig.SessMaxAge})
	session.Set("UserName", user.UserName)
	session.Set("Token", token)
	session.Set("IsLogin", true)
	// session.Set("UserLevel", user.UserLevel)
	session.Save()
}

/*
@Success 200 / return success
*/
func GetLogout(c *gin.Context) {
	destroySession(c)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// destroy session
func destroySession(c *gin.Context) {
	defer log.Debug("destroy session")
	session := sessions.Default(c)
	session.Options(sessions.Options{Path: config.AppConfig.SessPath, MaxAge: -1})
	session.Clear()
	session.Save()
}
