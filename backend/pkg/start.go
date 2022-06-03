package pkg

import (
	"gift/boot"
	"gift/lg"
	"github.com/gin-gonic/gin"
)

func Start() {

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	engine.GET("/ping", func(c *gin.Context) {
		boot.Sugar.Info("hello")
		lg.Info("hello")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.Run(":11111")
}
