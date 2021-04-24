// 管理员账户使用
package root

import (
	"course/handle"
	"course/handler"
	"course/request"
	"course/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
)

// 查询普通用户账户
func QueryUserCount(ctx *gin.Context) {
	content, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("read request body failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	req := &request.QueryUserCountReq{}
	err = json.Unmarshal(content, req)
	if err != nil {
		log.Println("Unmarshal content failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	data, err := handle.DoQueryUser(req)
	if err != nil {
		log.Println("doQuery user failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	response.SuccessResponse(ctx, data)
}


// 对用户账户进行操作: 修改 / 删除 / 禁用
func OperateUserCount(ctx *gin.Context) {
	action := ctx.Query(handler.Action)
	idStr := ctx.Param(handler.ParamID)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("parse id [ ", id, " ] failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	content, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("read request body failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	switch action {
	case handler.UpdateUserCount:
		err = handle.DoUpdateUser(int(id), content)
		if err != nil {
			response.FailedResponse(ctx, err)
			return
		}
		response.SuccessResponse(ctx, nil)
		return
	case handler.DeleteUserCount:
		err = handle.DoDeleteUser(int(id))
		if err != nil {
			response.FailedResponse(ctx, err)
			return
		}
		response.SuccessResponse(ctx, nil)
		return
	case handler.ChangeUserStatus:
		err = handle.DoChangeStatus(int(id), content)
		if err != nil {
			response.FailedResponse(ctx, err)
			return
		}
		response.SuccessResponse(ctx, nil)
		return
	default:
		response.FailedResponse(ctx, fmt.Errorf("action: %s 暂不支持", action))
		return
	}
}

// 添加普通用户账户
func AddUserCount(ctx *gin.Context) {
	_, isLogin := ctx.Get("user_name")
	if !isLogin {
		response.FailedResponse(ctx, fmt.Errorf("please login first"))
		return
	}
	// TODO: 只有管理员账户才能添加账户，这里需要对当前账户身份进行校验。
	content, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("read request body failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	req := &request.AddUserCountReq{}
	err = json.Unmarshal(content, req)
	if err != nil {
		log.Println("Unmarshal request failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	err = handle.DoAddUser(req)
	if err != nil {
		log.Println("add user [", req.UserName, " ] failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	response.SuccessResponse(ctx, nil)
}
