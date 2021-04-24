package handler

import (
	"course/dal"
	"course/handle"
	"course/request"
	"course/response"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

// 处理登录相关
func Login(ctx *gin.Context) {
	content, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		response.FailedResponse(ctx, err)
		log.Println("get Req Body failed: ", err.Error())
		return
	}
	req := &request.LoginReq{}
	err = json.Unmarshal(content, req)
	if err != nil {
		log.Println("Unmarshal LoginReq failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	user, err := handle.DoLogin(ctx, req)
	if err != nil {
		log.Println("login failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	session := sessions.Default(ctx) // 保存session
	session.Set("is_login", user.UserName)
	err = session.Save()

	if err != nil {
		log.Println("login failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	resp := &response.LoginResp{
		UserName: user.UserName,
		IsRoot:   req.IsRoot,
	}
	response.SuccessResponse(ctx, resp)
}
func LogOut(ctx *gin.Context)  {

}


func AddRootCount(ctx *gin.Context) {
	content, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		response.FailedResponse(ctx, err)
		return
	}
	req := &request.AddUserCountReq{}
	err = json.Unmarshal(content, req)
	if err != nil {
		response.FailedResponse(ctx, err)
		return
	}
	root := &dal.User{
		UserName:   req.UserName,
		PassWord:   req.PassWord,
		EmpType:    req.UserType,
		AuthLevel:  req.AuthLevel,
		Status:     req.Status,
		Department: req.Department,
		IsRoot:     1,
	}
	err = dal.AddUserCount(root)
	if err != nil {
		response.FailedResponse(ctx, err)
		return
	}
	response.SuccessResponse(ctx, nil)
}
