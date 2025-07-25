package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pento-zh/gin-test/cache"
)

func HelloRedis(c *gin.Context) {
	rdb := cache.GetRedisClient()
	res := rdb.HGetAll(context.Background(), "shoply_EJxsLN_user_1")

	fmt.Println("指令名称 ", res.Name())
	fmt.Println("指令全名 ", res.FullName())

	args := res.Args()
	for i := range args {
		fmt.Println("指令参数 ", args[i])
	}

	result, err := res.Result()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	fmt.Println(result)

	c.JSONP(http.StatusOK, gin.H{"name": res.Name(), "val": res.Val()})
}
