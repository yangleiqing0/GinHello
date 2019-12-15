package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"proxy_download/model"
	"regexp"
	"strconv"
	"strings"
)

func EmailDetail(context *gin.Context) {
	var email model.Email

	idString := context.Param("id")
	id, _ := strconv.Atoi(idString)

	emailDetail, err := email.Detail(id)
	if err != nil {
		fmt.Println("query table email err = ", err)
		return
	}
	context.JSON(http.StatusOK, Data{"list": emailDetail})
}

func EmailEdit(context *gin.Context) {
	var email model.Email

	if err := context.ShouldBind(&email); err != nil {
		context.JSON(http.StatusOK, Data{"err": "输入的数据不合法"})
		log.Panicln("err ->", err.Error())
		return
	}
	if email.ID != 0 {
		err := email.Update()
		if err != nil {
			fmt.Println("update email err = ", err)
			context.JSON(http.StatusBadRequest, Data{"err": "update email err" + err.Error()})
			return
		}
		context.JSON(http.StatusOK, Data{"msg": "update email success"})
		return
	}
	id, err := email.Save()
	if err != nil {
		fmt.Println("save email err ", err)
		context.JSON(http.StatusBadRequest, Data{"err": "save email err" + err.Error()})
		return
	}
	context.JSON(http.StatusOK, Data{"msg": "save email success, id:" + strconv.Itoa(id)})
}

func EmailList(context *gin.Context) {
	var email model.Email
	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))

	pagesize, err := strconv.Atoi(context.DefaultQuery("pagesize", "10"))

	fmt.Println("page ")
	emails, count, err := email.List(page, pagesize)
	if err != nil {
		err := fmt.Errorf("query table email err = %v", err.Error())
		fmt.Println(err)
		context.JSON(http.StatusBadGateway, err)
		return
	}
	context.JSON(http.StatusOK, Data{"list": emails, "count": count})
}

func EmailDel(context *gin.Context) {

	var emails NullMap
	var email model.Email
	err := context.BindJSON(&emails)

	if err != nil {
		log.Println("json.Unmarshal err = ", err)
		context.JSON(http.StatusOK, Data{"err": "get ids error"})
		return
	}

	switch ids := emails["ids"].(type) {
	// 对返回的元素进行判断   float64  id     []interface{} ids
	case float64:
		if err = email.Delete(ids); err != nil {
			fmt.Println("delete email err :", err)
			context.JSON(http.StatusBadRequest, Data{"err": err})
			return
		}
		context.JSON(http.StatusOK, Data{"msg": "del success"})
		return
	case []interface{}:
		if err = email.Deletes(ids); err != nil {
			fmt.Println("list delete email err :", err)
			context.JSON(http.StatusBadRequest, Data{"err": err})
			return
		}
		context.JSON(http.StatusOK, Data{"msg": "del list success"})
	}

}

func EmailNameValidate(context *gin.Context) {

	result, err := NameValidate(context, "emails")
	if err != nil {
		fmt.Println("err = ", err)
		context.JSON(http.StatusBadRequest, Data{"err": err.Error()})
		return
	}
	context.JSON(http.StatusOK, result)
}

func EmailToUserListValidate(context *gin.Context) {
	var params = struct {
		ToUserList string `json:"to_user_list"`
	}{}
	result := true

	err := context.BindJSON(&params)

	if err != nil {
		fmt.Println("context.BindJSON EmailToUserListValidate err = ", err)
		context.JSON(http.StatusBadRequest, Data{"err": err.Error()})
		return
	}

	UserList := strings.Split(params.ToUserList, ",")
	reg := regexp.MustCompile("^.+@(\\[?)[a-zA-Z0-9\\-.]+\\.([a-zA-Z]{2,3}|[0-9]{1,3})(]?)$")

	for i := 0; i < len(UserList); i++ {
		regResult := reg.FindAllStringSubmatch(UserList[i], -1)
		if len(regResult) == 0 {
			result = false
		}
	}

	context.JSON(http.StatusOK, result)
}
