package add

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Data struct {
	Num1   int `json:"num1" binding:"required"`
	Num2   int `json:"num2" binding:"required"`
	Result int
}

func Add(c *gin.Context) {
	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	data.Result = data.Num1 + data.Num2

	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{"result": data.Result})
}
