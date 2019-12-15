package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"proxy_download/model"
)

func NameValidate(context *gin.Context, tableName string) (result bool, err error) {

	var params = struct {
		Name   string `json:"name"`
		Id     int    `json:"id"`
		UserId int    `json:"user_id"`
	}{}

	err1 := context.BindJSON(&params)
	if err1 != nil {
		fmt.Println("context.BindJSON NameValidate err = ", err1)
		return
	}

	//name := strings.Trim(params.Name, " ")
	name := params.Name
	id := params.Id
	userId := params.UserId

	if id != 0 {
		result, err = model.UpdateNameValidate(name, tableName, id, userId)
		return
	}

	result, err = model.SaveNameValidate(name, tableName, userId)
	return
}
