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
	email := router.Group("/email")
	{
		email.GET("/detail/:id", handler.EmailDetail)
		email.GET("/list", handler.EmailList)
		email.POST("/del", handler.EmailDel)
		email.POST("/edit", handler.EmailEdit)
	}
	variable := router.Group("/variable")
	{
		variable.GET("/detail/:id", handler.VariableDetail)
		variable.GET("/list", handler.VariableList)
		variable.POST("/del", handler.VariableDel)
		variable.POST("/edit", handler.VariableEdit)
	}
	group := router.Group("/group")
	{
		group.GET("/detail/:id", handler.GroupDetail)
		group.GET("/list", handler.GroupList)
		group.POST("/del", handler.GroupDel)
		group.POST("/edit", handler.GroupEdit)
	}
	header := router.Group("/header")
	{
		header.GET("/detail/:id", handler.HeaderDetail)
		header.GET("/list", handler.HeaderList)
		header.POST("/del", handler.HeaderDel)
		header.POST("/edit", handler.HeaderEdit)
	}
	testCase := router.Group("/case")
	{
		testCase.GET("/detail/:id", handler.TestCaseDetail)
		testCase.GET("/list", handler.TestCaseList)
		testCase.POST("/del", handler.TestCaseDel)
		testCase.POST("/edit", handler.TestCaseEdit)
	}
	user := router.Group("/user")
	{
		user.GET("/detail/:id", handler.UserDetail)
		user.POST("/edit", handler.UserEdit)
	}

	return router
}
