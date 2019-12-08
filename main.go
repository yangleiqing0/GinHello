package main

import (
	"log"
	"proxy_download/initDB"
	"proxy_download/initRouter"
	"proxy_download/model"
)

// go mod init proxy_download  设置proxy_download来进行包管理

func main() {

	router := initRouter.SetupRouter()
	Db := initDB.DbInit()

	defer func() {
		err := Db.Close()
		if err != nil {
			log.Println("db.Close() err = ", err)
		}
	}()

	Db.AutoMigrate(&model.User{})

	_ = router.Run()

}
