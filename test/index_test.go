package test
import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"proxy_download/initRouter"
	"testing"
)
func TestIndexGetRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin", w.Body.String())
}
