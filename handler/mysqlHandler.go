package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proxy_download/initDB"
	"proxy_download/model"
)

var mysql, mysqls model.Mysql

func MysqlDetail(context *gin.Context) {
	db := initDB.DbInit()
	log.Println(">>> this is MysqlListHandler <<<")
	id := context.Param("id")
	log.Println("id = ", id)

	err := db.Where("id = ?", id).First(&model.Mysql{}).Scan(&mysql).Error
	if err != nil {
		fmt.Println("query table mysql err = ", err)
		return
	}
	log.Printf("mysql is %v", mysql.Ip)
	context.String(http.StatusOK, "hello "+mysql.Ip)
}

func MysqlEdit(context *gin.Context) {

	if err := context.ShouldBind(&mysql); err != nil {
		context.String(http.StatusBadRequest, "输入的数据不合法")
		log.Panicln("err ->", err.Error())
	}
	id, err := mysql.Save()
	if err != nil {
		fmt.Println("save mysql err ", err)
		context.String(http.StatusBadRequest, "save mysql err"+err.Error())

	}
	context.String(http.StatusOK, "编辑mysql成功, id:", id)
}

func MysqlList(context *gin.Context) {
	db := initDB.DbInit()
	r := map[string]interface{}{"list": nil}

	result := db.Find(&mysqls)
	if result.RecordNotFound() {
		context.JSON(http.StatusOK, r)
		return
	}

	err := result.Error

	if err != nil {
		err := fmt.Errorf("query table mysql err = %v", err.Error())
		fmt.Println(err)
		context.JSON(http.StatusBadGateway, err)
		return
	}
	fmt.Println("query result = ", mysqls)
	context.JSON(http.StatusOK, gin.H{"list": mysqls})
}
