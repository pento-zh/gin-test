package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/pento-zh/gin-test/db"
	"github.com/pento-zh/gin-test/db/model"
)

func ProductList(c *gin.Context) {
	db := db.GetDb()
	var products []model.Product
	db.Find(&products)
	c.JSON(200, products)
}

func ProductGet(c *gin.Context) {
	db := db.GetDb()
	// Read
	var product model.Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	c.JSON(200, product)
}

func ProductCreate(c *gin.Context) {
	db := db.GetDb()
	// Create
	pro := &model.Product{Code: "D42", Price: 100}
	if db.Create(pro).Error != nil {
		c.JSON(500, gin.H{"error": "Error creating product"})
		return
	}
	c.JSON(200, pro)
}

func ProductUpdate(c *gin.Context) {
	db := db.GetDb()

	var product model.Product
	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(model.Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
}

func ProductDelete(c *gin.Context) {
	db := db.GetDb()

	var product model.Product
	db.First(&product, 1) // 根据整型主键查找
	db.Delete(&product)
}
