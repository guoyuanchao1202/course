package normal

import (
	"course/handle"
	"course/handler"
	"course/request"
	"course/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
)

// 查询技术资料
func QueryTechniqueData(ctx *gin.Context) {
	content, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Read request Body failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}

	userName, _ := ctx.Get("user_name")
	req := &request.QueryDataForUserReq{}
	err = json.Unmarshal(content, req)
	if err != nil {
		log.Println("Unmarshal failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	res, err := handle.DoQueryDataByCond(userName.(string), req)
	if err != nil {
		log.Println("DoQueryDataByCond failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	response.SuccessResponse(ctx, res)
}

// 查看指定技术资料信息
func GetTechniqueDataInfo(ctx *gin.Context) {
	idStr := ctx.Param(handler.ParamID)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("parse id[ ", idStr, " ] failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	res, err := handle.DoQueryDataByID(id)
	if err != nil {
		log.Println("DoQueryDataByID failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	response.SuccessResponse(ctx, res)
}
