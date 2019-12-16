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
	id, _ := strconv.Atoi(idString)

	testCaseDetail, err := testCase.Detail(id)
	if err != nil {
		fmt.Println("query table testCase err = ", err)
		context.JSON(http.StatusBadRequest, Data{"list": testCaseDetail, "err": err.Error()})
		return
	}
	context.JSON(http.StatusOK, Data{"list": testCaseDetail})
}

func TestCaseEdit(context *gin.Context) {
	var testCase model.TestCase

	if err := context.ShouldBind(&testCase); err != nil {
		context.JSON(http.StatusOK, Data{"err": "输入的数据不合法"})
		log.Panicln("err ->", err.Error())
		return
	}
	if testCase.ID != 0 {
		err := testCase.Update()
		if err != nil {
			fmt.Println("update testCase err = ", err)
			context.JSON(http.StatusBadRequest, Data{"err": "update testCase err" + err.Error()})
			return
		}
		context.JSON(http.StatusOK, Data{"msg": "update testCase success"})
		return
	}
	id, err := testCase.Save()
	if err != nil {
		fmt.Println("save testCase err ", err)
		context.JSON(http.StatusBadRequest, Data{"err": "save testCase err" + err.Error()})
		return
	}
	context.JSON(http.StatusOK, Data{"msg": "save testCase success, id:" + strconv.Itoa(id)})
}

func TestCaseList(context *gin.Context) {
	var testCase model.TestCase

	var g model.Group
	var h model.Header
	groups, err := g.ListAll()
	headers, err := h.ListAll()

	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))

	pagesize, err := strconv.Atoi(context.DefaultQuery("pagesize", "10"))

	fmt.Println("page ")
	testCases, count, err := testCase.List(page, pagesize)
	if err != nil {
		err := fmt.Errorf("query table testCase err = %v", err.Error())
		fmt.Println(err)
		context.JSON(http.StatusBadGateway, err)
		return
	}
	context.JSON(http.StatusOK, Data{"list": testCases, "count": count, "groups": groups, "headers": headers})
}

func TestCaseDel(context *gin.Context) {

	var testCases NullMap
	var testCase model.TestCase
	err := context.BindJSON(&testCases)

	if err != nil {
		log.Println("json.Unmarshal err = ", err)
		context.JSON(http.StatusOK, Data{"err": "get ids error"})
		return
	}

	switch ids := testCases["ids"].(type) {
	// 对返回的元素进行判断   float64  id     []interface{} ids
	case float64:
		if err = testCase.Delete(ids); err != nil {
			fmt.Println("delete testCase err :", err)
			context.JSON(http.StatusBadRequest, Data{"err": err})
			return
		}
		context.JSON(http.StatusOK, Data{"msg": "del success"})
		return
	case []interface{}:
		if err = testCase.Deletes(ids); err != nil {
			fmt.Println("list delete testCase err :", err)
			context.JSON(http.StatusBadRequest, Data{"err": err})
			return
		}
		context.JSON(http.StatusOK, Data{"msg": "del list success"})
	}

}

func TestCaseNameValidate(context *gin.Context) {

	result, err := NameValidate(context, "test_cases")
	if err != nil {
		fmt.Println("err = ", err)
		context.JSON(http.StatusBadRequest, Data{"err": err.Error()})
		return
	}
	context.JSON(http.StatusOK, result)
}

func TestCaseRegularValidate(context *gin.Context) {
	var params = struct {
		Regular string `json:"regular"`
		UserId  int    `json:"user_id"`
	}{}

	err := context.BindJSON(&params)
	if err != nil {
		fmt.Println("TestCaseRegularValidate err = ", err)
		context.JSON(http.StatusOK, Data{"err": err.Error()})
		return
	}

	regular := params.Regular
	result := model.RegularValidate(regular)

	context.JSON(http.StatusOK, result)
}

func TestCaseHopeValidate(context *gin.Context) {
	var params = struct {
		HopeResult string `json:"hope_result"`
		UserId     int    `json:"user_id"`
	}{}

	err := context.BindJSON(&params)
	if err != nil {
		fmt.Println("TestCaseHopeValidate err = ", err)
		context.JSON(http.StatusOK, Data{"err": err.Error()})
		return
	}

	result := model.HopeResultValidate(params.HopeResult)
	context.JSON(http.StatusOK, result)
}

func TestCaseSqlVariableValidate(context *gin.Context) {

	var params = struct {
		Id          int    `json:"id"`
		SqlVariable string `json:"sql_variable"`
		UserId      int    `json:"user_id"`
	}{}

	err := context.BindJSON(&params)
	if err != nil {
		fmt.Println("TestCaseSqlVariableValidate err = ", err)
		context.JSON(http.StatusOK, Data{"err": err.Error()})
		return
	}

	if params.Id != 0 {
		result, err := model.UpdateSqlVariableValidate(params.SqlVariable, params.Id, params.UserId)
		if err != nil {
			fmt.Println("model.UpdateSqlVariableValidate err = ", err)
			context.JSON(http.StatusOK, Data{"err": err.Error()})
			return
		}
		context.JSON(http.StatusOK, result)
		return
	}

	result, err := model.CreateSqlVariableValidate(params.SqlVariable, params.UserId)
	if err != nil {
		fmt.Println("model.CreateSqlVariableValidate err = ", err)
		context.JSON(http.StatusOK, Data{"err": err.Error()})
		return
	}
	context.JSON(http.StatusOK, result)
}
