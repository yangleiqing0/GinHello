package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proxy_download/model"
	"strconv"
)

func TestCaseDetail(context *gin.Context) {
	var testCase model.TestCase

	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)

	testCaseDetail, err := testCase.Detail(id)
	if err != nil {
		fmt.Println("query table testCase err = ", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"list": testCaseDetail})
}

func TestCaseEdit(context *gin.Context) {
	var testCase model.TestCase

	if err := context.ShouldBind(&testCase); err != nil {
		context.JSON(http.StatusOK, gin.H{"err": "输入的数据不合法"})
		log.Panicln("err ->", err.Error())
		return
	}
	if testCase.ID != 0 {
		err := testCase.Update()
		if err != nil {
			fmt.Println("update testCase err = ", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": "update testCase err" + err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "update testCase success"})
		return
	}
	id, err := testCase.Save()
	if err != nil {
		fmt.Println("save testCase err ", err)
		context.JSON(http.StatusBadRequest, gin.H{"err": "save testCase err" + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "save testCase success, id:" + strconv.FormatInt(id, 10)})
}

func TestCaseList(context *gin.Context) {
	var testCase model.TestCase

	var g model.Group
	var h model.Header
	groups, err := g.ListAll()
	headers, err := h.ListAll()

	page, err := strconv.ParseInt(context.DefaultQuery("page", "1"), 10, 64)

	pagesize, err := strconv.ParseInt(context.DefaultQuery("pagesize", "10"), 10, 64)

	fmt.Println("page ")
	testCases, count, err := testCase.List(page, pagesize)
	if err != nil {
		err := fmt.Errorf("query table testCase err = %v", err.Error())
		fmt.Println(err)
		context.JSON(http.StatusBadGateway, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"list": testCases, "count": count, "groups": groups, "headers": headers})
}

func TestCaseDel(context *gin.Context) {

	var testCases NullMap
	var testCase model.TestCase
	err := context.BindJSON(&testCases)

	if err != nil {
		log.Println("json.Unmarshal err = ", err)
		context.JSON(http.StatusOK, gin.H{"err": "get ids error"})
		return
	}

	switch ids := testCases["ids"].(type) {
	// 对返回的元素进行判断   float64  id     []interface{} ids
	case float64:
		if err = testCase.Delete(ids); err != nil {
			fmt.Println("delete testCase err :", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "del success"})
		return
	case []interface{}:
		if err = testCase.Deletes(ids); err != nil {
			fmt.Println("list delete testCase err :", err)
			context.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "del list success"})
	}

}
