package main

import (
	"github.com/pento-zh/gin-test/router"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	router.Init(e)
	e.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
