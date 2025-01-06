package response

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, ResponseData{
		Code:    code,
		Message: GetResponseMessage(code),
		Data:    data,
	})
}

func ErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
