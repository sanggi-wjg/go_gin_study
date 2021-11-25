package main

import (
	"github.com/gin-gonic/gin"
	"github.nhnent.com/godo/cfo/models"
	"github.nhnent.com/godo/cfo/routers"
	"github.nhnent.com/godo/cfo/server/config"
	"github.nhnent.com/godo/cfo/server/log"
)

func init() {
	config.Init()
	models.Init()

	if config.ServerConfig.Debug {
		gin.ForceConsoleColor()
	}
	gin.SetMode(config.ServerConfig.RunMode)
	log.Init(config.ServerConfig.Debug)
}

func main() {
	router := routers.Init()
	err := router.Run(config.ServerConfig.HttpPort)
	if err != nil {
		panic(err)
	}
}
