package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	e.LoadHTMLGlob("templates/*")

	webGroup := e.Group("/")
	index(webGroup)
}

func index(g *gin.RouterGroup) {

	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello world 2025 !",
		})
	})
}
