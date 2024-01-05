package main

import (
	"fmt"

	"mid/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(func(ctx *gin.Context) {
		fmt.Println("你好,这是 middleware")
	})

	engine.GET("/redis", dao.GetFromRedis)
	engine.GET("/mysql", dao.GetFromMySQL)
	engine.GET("/openSearch", dao.GetFromOpenSearch)
	engine.GET("/common-openSearch", dao.GetFromCommonOpenSearch)

	dao.InitRedis()
	dao.InitMySQL()
	dao.InitOpenSearch()

	engine.Run(":9997")
}
