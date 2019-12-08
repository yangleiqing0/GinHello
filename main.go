package main

import (
	"proxy_download/initDB"
	"proxy_download/initRouter"
)

// go mod init proxy_download  设置proxy_download来进行包管理

func main() {

	router := initRouter.SetupRouter()

	initDB.DbInit()

	_ = router.Run()

}