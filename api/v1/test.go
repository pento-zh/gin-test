package v1

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the GORM API",
	})
}

func BindStruct(c *gin.Context) {
	type StructA struct {
		FieldA string `form:"field_a"`
	}
	var a StructA
	c.Bind(&a)

	c.JSON(200, gin.H{
		"a": a,
	})
}

func BindCheckout(c *gin.Context) {
	type Colors struct {
		Colors []string `form:"colors[]"`
	}
	var a Colors
	c.ShouldBind(&a)
	c.JSON(200, gin.H{
		"colors": a.Colors,
	})
}
