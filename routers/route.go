package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.nhnent.com/godo/cfo/middleware"
	"github.nhnent.com/godo/cfo/middleware/jwt"
	"github.nhnent.com/godo/cfo/routers/auth"
	"github.nhnent.com/godo/cfo/routers/health"
	"github.nhnent.com/godo/cfo/routers/v1/home"
	"github.nhnent.com/godo/cfo/server/config"
)

func Init() *gin.Engine {
	router := gin.New()
	// router.Use(server.Logger(), server.Recovery())

	store := cookie.NewStore([]byte(config.AppConfig.JwtSecret))
	router.Use(sessions.Sessions(config.AppConfig.SessPath, store))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Static(config.AppConfig.StaticURL, config.AppConfig.StaticRoot)
	router.Static(config.AppConfig.MediaURL, config.AppConfig.MediaRoot)
	router.StaticFile(config.AppConfig.FavIconURL, config.AppConfig.FavIconPath)
	router.LoadHTMLGlob(config.AppConfig.TemplateRoot)

	router.GET("/ping", health.Ping)

	router.GET("/login", auth.GetLogin)
	router.POST("/login", auth.PostLogin)
	router.GET("/logout", auth.GetLogout)

	authorized := router.Group("/")
	authorized.Use(middleware.Authorized())
	{
		authorized.GET("", home.Home)
		authorized.GET("/newuser", auth.GetNewUser)
		authorized.POST("/newuser", auth.PostNewUser)
	}

	jwt_v1 := router.Group("/v1")
	jwt_v1.Use(jwt.JWT())
	{

	}

	return router
}
