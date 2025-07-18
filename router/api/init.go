package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pento-zh/gin-test/api/v1"
)

func Init(e *gin.Engine) {
	apiGroup := e.Group("/api/v1")
	test(apiGroup)
}

func test(g *gin.RouterGroup) {
	testGroup := g.Group("/test")
	{
		testGroup.GET("/hello", v1.Hello)
		testGroup.GET("/bind", v1.BindStruct)
		testGroup.POST("/bind/checkout", v1.BindCheckout)
	}
}
