package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseData struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 响应消息
	Data    any    `json:"data"`    // 响应数据
}

func ParameterError(c *gin.Context) {
	code := http.StatusBadRequest
	c.JSON(code, responseData{
		Code:    code,
		Message: "参数错误",
		Data:    nil,
	})
}

func unauthorizedError(c *gin.Context) {
	code := http.StatusUnauthorized
	c.JSON(code, responseData{
		Code:    code,
		Message: "未经授权的",
		Data:    nil,
	})
}

func Custom(c *gin.Context, code int, message string, data any) {
	c.JSON(code, responseData{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func MissingParametersError(c *gin.Context) {
	code := http.StatusBadRequest
	c.JSON(code, responseData{
		Code:    code,
		Message: "缺少参数",
		Data:    nil,
	})
}

func Success(c *gin.Context, data any) {
	code := http.StatusOK
	c.JSON(code, responseData{
		Code:    code,
		Message: "成功",
		Data:    data,
	})
}

func Fail(c *gin.Context, message string) {
	code := http.StatusInternalServerError
	c.JSON(http.StatusInternalServerError, responseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
