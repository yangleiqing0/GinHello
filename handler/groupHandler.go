package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proxy_download/model"
	"strconv"
)

func GroupDetail(context *gin.Context) {
	var group model.Group

	idString := context.Param("id")
	id, _ := strconv.Atoi(idString)

	groupDetail, err := group.Detail(id)
	if err != nil {
		fmt.Println("query table group err = ", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"list": groupDetail})
}

func GroupEdit(context *gin.Context) {
	var group model.Group

	if err := context.ShouldBind(&group); err != nil {
		context.JSON(http.StatusOK, gin.H{"err": "输入的数据不合法"})
		log.Panicln("err ->", err.Error())
		return
	}
	if group.ID != 0 {
		err := group.Update()
		if err != nil {
			fmt.Println("update group err = ", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": "update group err" + err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "update group success"})
		return
	}
	id, err := group.Save()
	if err != nil {
		fmt.Println("save group err ", err)
		context.JSON(http.StatusBadRequest, gin.H{"err": "save group err" + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "save group success, id:" + strconv.FormatInt(id, 10)})
}

func GroupList(context *gin.Context) {
	var group model.Group
	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))

	pagesize, err := strconv.Atoi(context.DefaultQuery("pagesize", "10"))

	fmt.Println("page ")
	groups, count, err := group.List(page, pagesize)
	if err != nil {
		err := fmt.Errorf("query table group err = %v", err.Error())
		fmt.Println(err)
		context.JSON(http.StatusBadGateway, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"list": groups, "count": count})
}

func GroupDel(context *gin.Context) {

	var groups NullMap
	var group model.Group
	err := context.BindJSON(&groups)

	if err != nil {
		log.Println("json.Unmarshal err = ", err)
		context.JSON(http.StatusOK, gin.H{"err": "get ids error"})
		return
	}

	switch ids := groups["ids"].(type) {
	// 对返回的元素进行判断   float64  id     []interface{} ids
	case float64:
		if err = group.Delete(ids); err != nil {
			fmt.Println("delete group err :", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "del success"})
		return
	case []interface{}:
		if err = group.Deletes(ids); err != nil {
			fmt.Println("list delete group err :", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "del list success"})
	}

}

func GroupNameValidate(context *gin.Context) {

	result, err := NameValidate(context, "groups")
	if err != nil {
		fmt.Println("err = ", err)
		context.JSON(http.StatusBadRequest, Data{"err": err.Error()})
		return
	}
	context.JSON(http.StatusOK, result)
}
