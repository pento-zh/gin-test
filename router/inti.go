package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pento-zh/gin-test/router/api"
	"github.com/pento-zh/gin-test/router/web"
)

func Init(e *gin.Engine) {
	web.Init(e)
	api.Init(e)
}
