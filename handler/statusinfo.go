// 将一些全局变量映射关系传递过去，和前端的约定
package handler

import (
	"course/response"
	"course/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetStatusInfo(ctx *gin.Context) {
	_, isExist := ctx.Get("user_name")
	if !isExist {
		response.FailedResponse(ctx, fmt.Errorf("please login first"))
		return
	}
	res := &response.StatusInfoResp{
		CountStatusMap: utils.CountStatusMap,
		EmpTypeMap:     utils.EmpTypeMap,
		DataTypeMap:    utils.DataTypeMap,
		DepartmentMap:  utils.DepartmentMap,
	}
	response.SuccessResponse(ctx, res)
}
