package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proxy_download/model"
	"strconv"
)

func HeaderDetail(context *gin.Context) {
	var header model.Header

	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)

	headerDetail, err := header.Detail(id)
	if err != nil {
		fmt.Println("query table header err = ", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"list": headerDetail})
}

func HeaderEdit(context *gin.Context) {
	var header model.Header

	if err := context.ShouldBind(&header); err != nil {
		context.JSON(http.StatusOK, gin.H{"err": "输入的数据不合法"})
		log.Panicln("err ->", err.Error())
		return
	}
	if header.ID != 0 {
		err := header.Update()
		if err != nil {
			fmt.Println("update header err = ", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": "update header err" + err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "update header success"})
		return
	}
	id, err := header.Save()
	if err != nil {
		fmt.Println("save header err ", err)
		context.JSON(http.StatusBadRequest, gin.H{"err": "save header err" + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "save header success, id:" + strconv.FormatInt(id, 10)})
}

func HeaderList(context *gin.Context) {
	var header model.Header
	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))

	pagesize, err := strconv.Atoi(context.DefaultQuery("pagesize", "10"))
	fmt.Println("page ")
	headers, count, err := header.List(page, pagesize)
	if err != nil {
		err := fmt.Errorf("query table header err = %v", err.Error())
		fmt.Println(err)
		context.JSON(http.StatusBadGateway, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"list": headers, "count": count})
}

func HeaderDel(context *gin.Context) {

	var headers NullMap
	var header model.Header
	err := context.BindJSON(&headers)

	if err != nil {
		log.Println("json.Unmarshal err = ", err)
		context.JSON(http.StatusOK, gin.H{"err": "get ids error"})
		return
	}

	switch ids := headers["ids"].(type) {
	// 对返回的元素进行判断   float64  id     []interface{} ids
	case float64:
		if err = header.Delete(ids); err != nil {
			fmt.Println("delete header err :", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "del success"})
		return
	case []interface{}:
		if err = header.Deletes(ids); err != nil {
			fmt.Println("list delete header err :", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "del list success"})
	}

}
