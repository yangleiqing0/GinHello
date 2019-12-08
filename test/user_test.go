package test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"proxy_download/common"
	"proxy_download/initRouter"
	"testing"
)

var router *gin.Engine

func init() {
	router = initRouter.SetupRouter()
}

func TestUserPostForm(t *testing.T) {
	UserCommon("253775405@qq.com", "123456", common.RandString(10), 200, t)
}

func TestUserEditEmailError(t *testing.T) {
	UserCommon("yang", "123456", common.RandString(10), 400, t)
}

func TestUserEditNameNull(t *testing.T) {
	UserCommon("253775405@qq.com", "123456", "", 400, t)
}

func TestUserEditPasswordNull(t *testing.T) {
	UserCommon("253775405@qq.com", "", common.RandString(10), 400, t)
}

func TestUserEditEmailNull(t *testing.T) {
	UserCommon("", "123456", common.RandString(10), 400, t)
}

func UserCommon(email, password, name string, statusCode int, t *testing.T) {

	value := url.Values{}
	value.Add("email", email)
	value.Add("password", password)
	value.Add("name", name)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/edit", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, statusCode, w.Code)
}
