package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proxy_download/common"
	"proxy_download/model"
	"strconv"
	"strings"
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
	context.JSON(http.StatusOK, Data{"list": variableDetail})
}

func VariableEdit(context *gin.Context) {
	var variable model.Variable

	if err := context.ShouldBind(&variable); err != nil {
		context.JSON(http.StatusOK, Data{"err": "输入的数据不合法"})
		log.Panicln("err ->", err.Error())
		return
	}
	if variable.ID != 0 {
		err := variable.Update()
		if err != nil {
			fmt.Println("update variable err = ", err)
			context.JSON(http.StatusBadRequest, Data{"err": "update variable err" + err.Error()})
			return
		}
		context.JSON(http.StatusOK, Data{"msg": "update variable success"})
		return
	}
	id, err := variable.Save()
	if err != nil {
		fmt.Println("save variable err ", err)
		context.JSON(http.StatusBadRequest, Data{"err": "save variable err" + err.Error()})
		return
	}
	context.JSON(http.StatusOK, Data{"msg": "save variable success, id:" + strconv.FormatInt(id, 10)})
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
	context.JSON(http.StatusOK, Data{"list": variables, "count": count})
}

func VariableDel(context *gin.Context) {

	var variables NullMap
	var variable model.Variable
	err := context.BindJSON(&variables)

	if err != nil {
		log.Println("json.Unmarshal err = ", err)
		context.JSON(http.StatusOK, Data{"err": "get ids error"})
		return
	}

	switch ids := variables["ids"].(type) {
	// 对返回的元素进行判断   float64  id     []interface{} ids
	case float64:
		if err = variable.Delete(ids); err != nil {
			fmt.Println("delete variable err :", err)
			context.JSON(http.StatusBadRequest, Data{"err": err})
			return
		}
		context.JSON(http.StatusOK, Data{"msg": "del success"})
		return
	case []interface{}:
		if err = variable.Deletes(ids); err != nil {
			fmt.Println("list delete variable err :", err)
			context.JSON(http.StatusBadRequest, Data{"err": err})
			return
		}
		context.JSON(http.StatusOK, Data{"msg": "del list success"})
	}

}

func VariableNameValidate(context *gin.Context) {

	var params = struct {
		Name             string `json:"name"`
		RegisterVariable string `json:"register_variable"`
		VariableId       int    `json:"variable_id"`
		TestcaseId       int    `json:"testcase_id"`
		UserId           int    `json:"user_id"`
		Update           bool   `json:"update"`
	}{}

	err1 := context.BindJSON(&params)
	if err1 != nil {
		fmt.Println("context.BindJSON VariableNameValidate err = ", err1)
		return
	}

	update := params.Update
	variableId := params.VariableId
	userId := params.UserId
	name := params.Name
	testcaseId := params.TestcaseId
	registerVariable := params.RegisterVariable

	if update {
		if variableId != 0 {
			result, err := model.UpdateVariableNameValidate(variableId, userId, name)
			if err != nil {
				context.JSON(http.StatusBadRequest, Data{"err": err.Error()})
				return
			}
			context.JSON(http.StatusOK, result)
			return
		}
		result, err := model.UpdateCaseRegisterNameValidate(testcaseId, userId, registerVariable)
		if err != nil {
			context.JSON(http.StatusBadRequest, Data{"err": err.Error()})
			return
		}
		context.JSON(http.StatusOK, result)
		return
	}

	if name != "" {
		registerVariable = name
	}

	if strings.Index(registerVariable, ",") != -1 && registerVariable != "" && len(strings.Trim(registerVariable, " ")) > 0 {
		variableList := strings.Split(registerVariable, ",")
		if len(variableList) != len(common.SliceToMap(variableList)) {
			context.JSON(http.StatusOK, false)
			return
		}

		for _, vbName := range variableList {
			count, _ := model.QueryVariableCount(vbName, 0, userId)
			if count != 0 {
				context.JSON(http.StatusOK, false)
				return
			}
		}
		context.JSON(http.StatusOK, true)
		return
	}

	count, _ := model.QueryVariableCount(registerVariable, 0, userId)

	context.JSON(http.StatusOK, count == 0)
}
