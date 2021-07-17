package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	{
		"code":10001, //错误码
		"msg":"XXX",  //提示信息
		"data":{}     //数据
	}
*/
type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseSuccess(c *gin.Context) {

	var res = new(ResponseData)
	res.Code = CodeSuccess
	res.Msg = CodeSuccess.Msg()
	c.JSON(http.StatusOK, res)
}
func ResponseSuccessWithCode(c *gin.Context, code ResCode) {

	var res = new(ResponseData)
	res.Code = code
	res.Msg = code.Msg()
	c.JSON(http.StatusOK, res)
}

func ResponseSuccessWithMsg(c *gin.Context, code ResCode, msg string) {

	var res = new(ResponseData)
	res.Code = code
	res.Msg = msg
	c.JSON(http.StatusOK, res)
}

func ResponseSuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
	})
}
