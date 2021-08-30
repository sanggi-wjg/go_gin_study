package routers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", Ping)
	router.GET("/users", Users)

	// 	user := routers.Group("/user")
	// 	{
	// 		// This handler will match /user/john but will not match /user/ or /user
	// 		user.GET("/user/:name", routers.HelloName)
	// 		// However, this one will match /user/john/ and also /user/john/send
	// 		// If no other routers match /user/john, it will redirect to /user/john/
	// 		user.GET("/user/:name/*action", routers.HelloNameAction)
	// 	}

	// 	// authorized := routers.Group("/")
	// 	// authorized.Use(AuthRequired()){
	// 	// }
}
