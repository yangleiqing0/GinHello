package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proxy_download/model"
	"strconv"
)

type NullMap map[string]interface{}

func MysqlDetail(context *gin.Context) {
	var mysql model.Mysql

	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)

	mysqlDetail, err := mysql.Detail(id)
	if err != nil {
		fmt.Println("query table mysql err = ", err)
		return
	}
	log.Printf("mysql is %v", mysqlDetail.Ip)
	context.JSON(http.StatusOK, gin.H{"list": mysqlDetail})
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
	var mysql model.Mysql
	page, err := strconv.ParseInt(context.DefaultQuery("page", "1"), 10, 64)

	pagesize, err := strconv.ParseInt(context.DefaultQuery("pagesize", "10"), 10, 64)

	fmt.Println("page ")
	mysqls, count, err := mysql.List(page, pagesize)
	if err != nil {
		err := fmt.Errorf("query table mysql err = %v", err.Error())
		fmt.Println(err)
		context.JSON(http.StatusBadGateway, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"list": mysqls, "count": count})
}

func MysqlDel(context *gin.Context) {

	var mysqls NullMap
	var mysql model.Mysql
	err := context.BindJSON(&mysqls)

	if err != nil {
		log.Println("json.Unmarshal err = ", err)
		context.JSON(http.StatusOK, gin.H{"err": "get ids error"})
		return
	}

	switch ids := mysqls["ids"].(type) {
	// 对返回的元素进行判断   float64  id     []interface{} ids
	case float64:
		if err = mysql.Delete(ids); err != nil {
			fmt.Println("delete mysql err :", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "del success"})
		return
	case []interface{}:
		if err = mysql.Deletes(ids); err != nil {
			fmt.Println("list delete mysql err :", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "del list success"})
	}

}
