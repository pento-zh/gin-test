package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pento-zh/gin-test/api/v1"
)

func Init(e *gin.Engine) {
	apiGroup := e.Group("/api/v1")
	test(apiGroup)
	redis(apiGroup)
	db(apiGroup)
}

func test(g *gin.RouterGroup) {
	testGroup := g.Group("/test")
	{
		testGroup.GET("/hello", v1.Hello)
		testGroup.GET("/bind", v1.BindStruct)
		testGroup.POST("/bind/checkout", v1.BindCheckout)
	}
}

func redis(g *gin.RouterGroup) {
	redisGroup := g.Group("/redis")
	{
		redisGroup.GET("/hello", v1.HelloRedis)
	}
}

func db(g *gin.RouterGroup) {
	dbGroup := g.Group("/db")
	{
		dbGroup.GET("/product/list", v1.ProductList)
		dbGroup.POST("/product/create", v1.ProductCreate)
		dbGroup.GET("/product/:id", v1.ProductGet)
		dbGroup.PUT("/product/:id", v1.ProductUpdate)
		dbGroup.DELETE("/product/:id", v1.ProductDelete)
	}
}
