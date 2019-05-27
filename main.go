package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"
	"github.com/wylzabc/jarvis/add"
	//	"strconv"
)

var router *gin.Engine

func InitRouter() {
	//	gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router = gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/add/add", add.Add)
}
func main() {
	InitRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
