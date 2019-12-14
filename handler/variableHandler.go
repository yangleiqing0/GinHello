package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proxy_download/model"
	"strconv"
)

func VariableDetail(context *gin.Context) {
	var variable model.Variable

	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)

	variableDetail, err := variable.Detail(id)
	if err != nil {
		fmt.Println("query table variable err = ", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"list": variableDetail})
}

func VariableEdit(context *gin.Context) {
	var variable model.Variable

	if err := context.ShouldBind(&variable); err != nil {
		context.JSON(http.StatusOK, gin.H{"err": "输入的数据不合法"})
		log.Panicln("err ->", err.Error())
		return
	}
	if variable.ID != 0 {
		err := variable.Update()
		if err != nil {
			fmt.Println("update variable err = ", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": "update variable err" + err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "update variable success"})
		return
	}
	id, err := variable.Save()
	if err != nil {
		fmt.Println("save variable err ", err)
		context.JSON(http.StatusBadRequest, gin.H{"err": "save variable err" + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "save variable success, id:" + strconv.FormatInt(id, 10)})
}

func VariableList(context *gin.Context) {
	var variable model.Variable
	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))

	pagesize, err := strconv.Atoi(context.DefaultQuery("pagesize", "10"))

	fmt.Println("page ")
	variables, count, err := variable.List(page, pagesize)
	if err != nil {
		err := fmt.Errorf("query table variable err = %v", err.Error())
		fmt.Println(err)
		context.JSON(http.StatusBadGateway, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"list": variables, "count": count})
}

func VariableDel(context *gin.Context) {

	var variables NullMap
	var variable model.Variable
	err := context.BindJSON(&variables)

	if err != nil {
		log.Println("json.Unmarshal err = ", err)
		context.JSON(http.StatusOK, gin.H{"err": "get ids error"})
		return
	}

	switch ids := variables["ids"].(type) {
	// 对返回的元素进行判断   float64  id     []interface{} ids
	case float64:
		if err = variable.Delete(ids); err != nil {
			fmt.Println("delete variable err :", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "del success"})
		return
	case []interface{}:
		if err = variable.Deletes(ids); err != nil {
			fmt.Println("list delete variable err :", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "del list success"})
	}

}
