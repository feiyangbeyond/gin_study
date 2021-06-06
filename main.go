package main

import (
	"gin_study/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	httpPort int
	wsPort   int
)

func init() {
	httpPort, wsPort, _ = util.GetServer()
}

func main() {

	gin.ForceConsoleColor()

	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		util.LogInfo(c.FullPath())
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	r.Run(":" + strconv.Itoa(httpPort))
}
