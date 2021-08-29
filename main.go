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

	if server.ServerConfig.Debug {
		gin.ForceConsoleColor()
	}
	gin.SetMode(server.ServerConfig.RunMode)
}

func main() {
	router := gin.New()
	router.Use(server.Logger(), server.Recovery())

	router.Static(server.AppConfig.StaticURL, server.AppConfig.StaticRoot)
	router.Static(server.AppConfig.MediaURL, server.AppConfig.MediaRoot)
	router.StaticFile(server.AppConfig.FavIconURL, server.AppConfig.FavIconPath)
	// router.LoadHTMLGlob(server.AppConfig.TemplateRoot)

	server.RegisterRoutes(router)

	err := router.Run(server.ServerConfig.HttpPort)
	if err != nil {
		panic(err)
	}
}
