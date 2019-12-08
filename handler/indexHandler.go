package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func IndexHandler(context *gin.Context) {
	log.Println(">>> this is login <<<")
	context.String(http.StatusOK, "hello gin")
}
