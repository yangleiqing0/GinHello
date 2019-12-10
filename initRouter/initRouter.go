package initRouter

import (
	"github.com/gin-gonic/gin"
	"proxy_download/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 添加 Get 请求路由
	router.GET("/", handler.IndexHandler)

	mysql := router.Group("/mysql")
	{
		mysql.GET("/detail/:id", handler.MysqlDetail)
		mysql.GET("/list", handler.MysqlList)
		mysql.POST("/del", handler.MysqlDel)
		mysql.POST("/edit", handler.MysqlEdit)
	}
	user := router.Group("/user")
	{
		user.GET("/detail/:id", handler.UserDetail)
		user.POST("/edit", handler.UserEdit)
	}

	return router
}
