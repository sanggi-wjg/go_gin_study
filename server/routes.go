package server

import (
	"go_gin_study/home"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", home.Ping)

	// 	user := router.Group("/user")
	// 	{
	// 		// This handler will match /user/john but will not match /user/ or /user
	// 		user.GET("/user/:name", home.HelloName)
	// 		// However, this one will match /user/john/ and also /user/john/send
	// 		// If no other routers match /user/john, it will redirect to /user/john/
	// 		user.GET("/user/:name/*action", home.HelloNameAction)
	// 	}

	// 	// authorized := router.Group("/")
	// 	// authorized.Use(AuthRequired()){
	// 	// }
}
