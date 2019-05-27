package multiplication

import (
	//	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wylzabc/jarvis/util"
	//"strconv"
	"encoding/json"
	"github.com/wylzabc/jarvis/data"
	"testing"
)

var router *gin.Engine

func Init() {
	gin.SetMode(gin.DebugMode)
	router = gin.Default()
	router.POST("/multi/multi", Multi)
}

//type AddData struct {
//	A      int
//	B      int
//	Sum    int
//	AbsSum int
//}
//
//var addDatas = []AddData{
//	{1, 2, 3, 3},
//	{1, -2, -1, 3},
//	{-1, 2, 1, 3},
//	{-1, -2, -3, 3},
//}
//

func TestAdd(t *testing.T) {
	Init()
	uri := "/multi/multi"
	param := make(map[string]interface{})
	param["num1"] = 1
	param["num2"] = 2

	body := util.PostJson(uri, param, router)
	var result data.Result
	err := json.Unmarshal(body, &result)
	if err != nil {
		t.Errorf("解析响应出错，err:%v\n", err)
	}
	if result.Result != 2 {
		t.Errorf("响应数据不符，result: %v\n", result)
	}
}

//func TestAbsAdd(t *testing.T) {
//	Init()
//	for _, testData := range addDatas {
//		uri := "/absadd/"
//		uri += strconv.Itoa(testData.A) + "/" + strconv.Itoa(testData.B)
//		body := util.Get(uri, router)
//		sum, _ := strconv.Atoi(string(body))
//		if sum != testData.AbsSum {
//			t.Errorf("error. body: %d", sum)
//		}
//	}
//
//}
