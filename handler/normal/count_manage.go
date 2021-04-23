package normal

import (
	"course/handle"
	"course/request"
	"course/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// 修改密码
func ChangePassWord(ctx *gin.Context) {
	userName, isExist := ctx.Get("user_name")
	if !isExist {
		response.FailedResponse(ctx, fmt.Errorf("user_name is not allowed nil"))
		return
	}
	content, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		response.FailedResponse(ctx, err)
		return
	}
	req := &request.ChangePassWordForUser{}
	err = json.Unmarshal(content, req)
	if err != nil {
		response.FailedResponse(ctx, err)
		return
	}
	err = handle.DoChangePW(userName.(string), req)
	if err != nil {
		response.FailedResponse(ctx, err)
		return
	}
	response.SuccessResponse(ctx, nil)
}
