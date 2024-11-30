package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bytedance/arishem/arishem"
	"github.com/gin-gonic/gin"
)

func ArishemPost(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}
	// 重置请求体，以便后续中间件可以再次读取
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	// 将请求体中的JSON数据转换成字符串，并格式化输出
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "    ") // 使用四个空格进行缩进
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to format JSON"})
		return
	}
	condition := prettyJSON.String()
	pass, err := arishem.JudgeCondition(c, condition)
    if err != nil {
        println(err.Error())
    }
    if pass {
        println("condition passed!")
    }

	c.JSON(http.StatusOK, gin.H{"received_condition": condition})
}