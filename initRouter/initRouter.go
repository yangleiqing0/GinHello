package initRouter

import (
	"github.com/gin-gonic/gin"
	"proxy_download/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 添加 Get 请求路由
	router.GET("/", handler.IndexHandler)

	user := router.Group("/user")
	{
		user.GET("/:name", handler.UserList)
		user.POST("/edit", handler.UserEdit)
	}


	return router
}