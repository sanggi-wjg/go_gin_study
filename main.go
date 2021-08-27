package main

import (
	"go_gin_study/server"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(server.Logger(), server.Recovery())

	config := server.NewConfig()
	settings := config.GetConfigSettings()

	router.Static(settings.StaticURL, settings.StaticRoot)
	router.Static(settings.MediaURL, settings.MediaRoot)
	router.StaticFile(settings.FavIconURL, settings.FavIconPath)
	// router.LoadHTMLGlob(settings.TemplateRoot)

	server.RegisterRoutes(router, config)

	router.Run(settings.Port)
}
