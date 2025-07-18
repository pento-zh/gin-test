package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pento-zh/gin-test/router/api/v1"
)

func Init(e *gin.Engine) {

	apiGroup := e.Group("/api")
	v1.Init(apiGroup)
}
