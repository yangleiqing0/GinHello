package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proxy_download/initDB"
	"proxy_download/model"
	"strconv"
)

type NullInterface interface {
}

func MysqlDetail(context *gin.Context) {
	var mysql model.Mysql
	db := initDB.DbInit()
	log.Println(">>> this is MysqlDetailHandler <<<")
	id := context.Param("id")
	log.Println("id = ", id)

	err := db.Where("id = ?", id).First(&model.Mysql{}).Scan(&mysql).Error
	if err != nil {
		fmt.Println("query table mysql err = ", err)
		return
	}
	log.Printf("mysql is %v", mysql.Ip)
	context.JSON(http.StatusOK, gin.H{"list": mysql})
}

func MysqlEdit(context *gin.Context) {
	var mysql model.Mysql

	if err := context.ShouldBind(&mysql); err != nil {
		context.JSON(http.StatusOK, gin.H{"err": "输入的数据不合法"})
		log.Panicln("err ->", err.Error())
		return
	}
	if mysql.ID != 0 {
		err := mysql.Update()
		if err != nil {
			fmt.Println("update mysql err = ", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": "update mysql err" + err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "update mysql success"})
		return
	}
	id, err := mysql.Save()
	if err != nil {
		fmt.Println("save mysql err ", err)
		context.JSON(http.StatusBadRequest, gin.H{"err": "save mysql err" + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "save mysql success, id:" + strconv.FormatInt(id, 10)})
}

func MysqlList(context *gin.Context) {
	var mysqls []model.Mysql
	db := initDB.DbInit()

	err := db.Order("updated_at desc, id desc").Find(&mysqls).Error

	if err != nil {
		err := fmt.Errorf("query table mysql err = %v", err.Error())
		fmt.Println(err)
		context.JSON(http.StatusBadGateway, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"list": mysqls})
}

func MysqlDel(context *gin.Context) {

	var mysqls NullInterface

	err := context.BindJSON(&mysqls)

	if err != nil {
		log.Println("json.Unmarshal err = ", err)
		context.JSON(http.StatusOK, gin.H{"err": "get ids error"})
		return
	}
	if ids, ok := mysqls.(map[string]interface{}); ok == true {
		log.Println(ids)
	}
	context.JSON(http.StatusOK, gin.H{"msg": "del success"})

}
