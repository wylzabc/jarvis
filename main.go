package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"
	"strconv"
)

var router *gin.Engine

func add(c *gin.Context) {
	a, _ := strconv.Atoi(c.Param("a"))
	b, _ := strconv.Atoi(c.Param("b"))
	c.String(200, "the result:%d\n", a+b)
}

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)
	router = gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/add/:a/:b", add)
}
func main() {
	InitRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
