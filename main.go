package main

import (
	"go_gin_study/models"
	"go_gin_study/server"

	"github.com/gin-gonic/gin"
)

func init() {
	server.InitConfig()
	server.InitLogger(server.ServerConfig.Debug)
	models.InitDatabase()
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
