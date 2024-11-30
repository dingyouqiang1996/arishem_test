package controllers

import "github.com/gin-gonic/gin"

type ArishemForm struct {
	Number1 int `form:"number1"`
	Expr string `form:"expr"`
	Number2 int `form:"number2"`
	Operator string `form:"operator"`
	Result int `form:"result"`
}

func ArishemPost(c *gin.Context) {

}
