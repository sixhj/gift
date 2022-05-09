package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.POST("/hello", func(context *gin.Context) {
		context.JSON(200, "hello")
		return
	})
	engine.Run()
}
