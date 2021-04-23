// 成功 / 失败 通用返回方法
package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode    = 0
	FailedCode     = -1
	SuccessMessage = "success"
)

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	resp := &CommonResponse{
		Code:    SuccessCode,
		Message: SuccessMessage,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, resp)
}

func FailedResponse(ctx *gin.Context, err error) {
	resp := &CommonResponse{
		Code:    FailedCode,
		Message: err.Error(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, resp)
}
