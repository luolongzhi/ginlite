package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//错误码格式定义
type ErrorBody struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Local     string `json:"local"`
}

var ERRMSG_MAP map[string]ErrorBody = map[string]ErrorBody{
	"Error.Common.NotFound":   ErrorBody{ErrorCode: 1000 + 1, Message: "record not found", Local: "记录未找到"},
	"Error.Common.ParamWrong": ErrorBody{ErrorCode: 1000 + 2, Message: "lack of parameter or wrong", Local: "参数缺失或错误"}}

//统一返回格式定义
type ResponseFormat struct {
	RequestId string      `json:"request_id"`
	ErrorCode int         `json:"error_code"`
	Response  interface{} `json:"response"`
}

func ResJson(c *gin.Context, data interface{}) {
	requestId := c.MustGet("_self_context__request_id").(string)

	c.JSON(200, ResponseFormat{RequestId: requestId, ErrorCode: 0, Response: data})
}

func ResError(c *gin.Context, errMsg string) {
	requestId := c.MustGet("_self_context__request_id").(string)

	data, ok := ERRMSG_MAP[errMsg]
	fmt.Println("000000000000 %v %v %v", data, ok, errMsg)

	if ok {
		c.JSON(400, ResponseFormat{RequestId: requestId, ErrorCode: data.ErrorCode, Response: data})
	} else {
		localErr := "未知系统错误"
		body := ErrorBody{
			ErrorCode: -1,
			Message:   errMsg,
			Local:     localErr}
		c.JSON(400, ResponseFormat{RequestId: requestId, ErrorCode: -1, Response: body})
	}
}
