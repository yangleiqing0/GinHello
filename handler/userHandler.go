package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proxy_download/initDB"
	"proxy_download/model"
)

func UserDetail(context *gin.Context) {
	db := initDB.DbInit()
	log.Println(">>> this is UserListHandler <<<")
	id := context.Param("id")
	log.Println("id = ", id)
	var user model.User
	err := db.Where("id = ?", id).First(&model.User{}).Scan(&user).Error
	if err != nil {
		fmt.Println("query user err = ", err)
		return
	}
	log.Printf("user is %v", user.Name)
	context.String(http.StatusOK, "hello "+user.Name)
}

func UserEdit(context *gin.Context) {

	var user model.User
	if err := context.ShouldBind(&user); err != nil {
		context.String(http.StatusBadRequest, "输入的数据不合法")
		log.Panicln("err ->", err.Error())
	}
	id, err := user.Save()
	if err != nil {
		fmt.Println("save user err ", err)
		context.String(http.StatusBadRequest, "save user err"+err.Error())

	}
	context.String(http.StatusOK, "编辑用户成功, id:", id)
}
