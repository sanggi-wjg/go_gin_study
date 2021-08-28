// package main

// import (
// 	"go_gin_study/server"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// Config
// 	config := server.NewConfig()
// 	settings := config.Settings()

// 	if settings.Debug == true {
// 		gin.SetMode(gin.DebugMode)
// 	} else {
// 		gin.SetMode(gin.ReleaseMode)
// 	}

// 	// gin
// 	router := gin.New()
// 	router.Use(server.Logger(), server.Recovery())

// 	// Route
// 	router.Static(settings.StaticURL, settings.StaticRoot)
// 	router.Static(settings.MediaURL, settings.MediaRoot)
// 	router.StaticFile(settings.FavIconURL, settings.FavIconPath)
// 	// router.LoadHTMLGlob(settings.TemplateRoot)

// 	server.RegisterRoutes(router, config)

// 	// Run
// 	router.Run(settings.Port)
// }

package main

import (
	"go_gin_study/server"

	"github.com/gin-gonic/gin"
)

func init() {
	server.InitConfig()
	server.InitLogger(server.ServerConfig.Debug)
}

func main() {
	gin.SetMode(server.ServerConfig.RunMode)

	router := gin.New()
	router.Use(server.Logger(), server.Recovery())

	router.Static(server.AppConfig.StaticURL, server.AppConfig.StaticRoot)
	router.Static(server.AppConfig.MediaURL, server.AppConfig.MediaRoot)
	router.StaticFile(server.AppConfig.FavIconURL, server.AppConfig.FavIconPath)
	// router.LoadHTMLGlob(server.AppConfig.TemplateRoot)

	server.RegisterRoutes(router)

	router.Run(server.ServerConfig.HttpPort)
}
