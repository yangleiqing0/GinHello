package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proxy_download/model"
	"strconv"
)

func EmailDetail(context *gin.Context) {
	var email model.Email

	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)

	emailDetail, err := email.Detail(id)
	if err != nil {
		fmt.Println("query table email err = ", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"list": emailDetail})
}

func EmailEdit(context *gin.Context) {
	var email model.Email

	if err := context.ShouldBind(&email); err != nil {
		context.JSON(http.StatusOK, gin.H{"err": "输入的数据不合法"})
		log.Panicln("err ->", err.Error())
		return
	}
	if email.ID != 0 {
		err := email.Update()
		if err != nil {
			fmt.Println("update email err = ", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": "update email err" + err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "update email success"})
		return
	}
	id, err := email.Save()
	if err != nil {
		fmt.Println("save email err ", err)
		context.JSON(http.StatusBadRequest, gin.H{"err": "save email err" + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "save email success, id:" + strconv.FormatInt(id, 10)})
}

func EmailList(context *gin.Context) {
	var email model.Email
	page, err := strconv.ParseInt(context.DefaultQuery("page", "1"), 10, 64)

	pagesize, err := strconv.ParseInt(context.DefaultQuery("pagesize", "10"), 10, 64)

	fmt.Println("page ")
	emails, count, err := email.List(page, pagesize)
	if err != nil {
		err := fmt.Errorf("query table email err = %v", err.Error())
		fmt.Println(err)
		context.JSON(http.StatusBadGateway, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"list": emails, "count": count})
}

func EmailDel(context *gin.Context) {

	var emails NullMap
	var email model.Email
	err := context.BindJSON(&emails)

	if err != nil {
		log.Println("json.Unmarshal err = ", err)
		context.JSON(http.StatusOK, gin.H{"err": "get ids error"})
		return
	}

	switch ids := emails["ids"].(type) {
	// 对返回的元素进行判断   float64  id     []interface{} ids
	case float64:
		if err = email.Delete(ids); err != nil {
			fmt.Println("delete email err :", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "del success"})
		return
	case []interface{}:
		if err = email.Deletes(ids); err != nil {
			fmt.Println("list delete email err :", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "del list success"})
	}

}
