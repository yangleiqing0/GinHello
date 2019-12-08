package handler


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proxy_download/model"
)

func UserList(context *gin.Context) {
	log.Println(">>> this is UserListHandler <<<")
	username := context.Param("name")
	age := context.Query("age")
	context.String(http.StatusOK, "hello" +  username +  age + "岁 ")
}

func UserEdit(context *gin.Context) {
	//name := context.PostForm("name")
	//email := context.PostForm("email")
	//password := context.DefaultPostForm("password", "Wa123456")
	//println("name:", name, "email:", email, "password:", password)
	var user model.User
	if err := context.ShouldBind(&user); err != nil {
		context.String(http.StatusBadRequest, "输入的数据不合法")
		log.Panicln("err ->", err.Error())
	}
	id, err := user.Save()
	if err != nil{
	    fmt.Println("save user err ", err)
	    return
	}

	context.String(http.StatusOK, "编辑用户成功, id:", id)
}
