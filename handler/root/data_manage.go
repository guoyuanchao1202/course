// 管理员账户使用
package root

import (
	"course/dal"
	"course/handle"
	"course/handler"
	"course/request"
	"course/response"
	"course/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"mime/multipart"
	"strconv"
)

// 查询技术资料
func QueryTechniqueData(ctx *gin.Context) {
	content, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("read request body failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	req := &request.QueryDataForRootReq{}
	err = json.Unmarshal(content, req)
	if err != nil {
		log.Println("Unmarshal QueryDataForRootReq failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	res, err := handle.DoQueryData(req)
	if err != nil {
		log.Println("DoQueryData failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	response.SuccessResponse(ctx, res)
}

// 添加技术资料
func AddTechniqueData(ctx *gin.Context) {
	userName, _ := ctx.Get("user_name")
	// 获取form表单内容
	multi, err := ctx.MultipartForm()
	if err != nil {
		response.FailedResponse(ctx, err)
		return
	}
	req, err := getAddDataForRootReq(multi)
	if err != nil {
		response.FailedResponse(ctx, err)
		return
	}
	upLoadFile := multi.File["upload"]
	err = handle.DoAddData(userName.(string), req, upLoadFile[0])
	if err != nil {
		log.Println("DoAddData failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	response.SuccessResponse(ctx, nil)
}

func getAddDataForRootReq(multiForm *multipart.Form) (*request.AddDataForRootReq, error) {
	title := multiForm.Value["title"]
	res := &request.AddDataForRootReq{}
	if len(title) != 1 {
		return nil, fmt.Errorf("expect len(title) == 1, but len(title) == %s", len(title))
	}
	res.Title = title[0]
	intro := multiForm.Value["introduction"]
	if len(title) != 1 {
		return nil, fmt.Errorf("expect len(intro) == 1, but len(intro) == %s", len(intro))
	}
	res.Introduction = intro[0]
	dataType := multiForm.Value["data_type"]
	if len(dataType) != 1 {
		return nil, fmt.Errorf("expect len(dataType) == 1, but len(dataType) == %s", len(dataType))
	}
	dataTypeCode, err := strconv.Atoi(dataType[0])
	if err != nil {
		return nil, err
	}
	res.DataType = dataTypeCode
	authLevel := multiForm.Value["auth_level"]
	if len(authLevel) != 1 {
		return nil, fmt.Errorf("expect len(authLevel) == 1, but len(authLevel) == %s", len(authLevel))
	}
	authLevelCode, err := strconv.Atoi(authLevel[0])
	if err != nil {
		return nil, err
	}
	res.AuthLevel = authLevelCode

	department := multiForm.Value["department"]
	if len(department) != 1 {
		return nil, fmt.Errorf("expect len(department) == 1, but len(department) == %s", len(department))
	}
	departmentCode, err := strconv.Atoi(department[0])
	if err != nil {
		return nil, err
	}
	res.Department = departmentCode
	return res, nil
}

// 获取指定技术资料的详细信息
func GetTechniqueDataInfo(ctx *gin.Context) {
	idStr := ctx.Param(handler.ParamID)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.FailedResponse(ctx, err)
		return
	}
	action := ctx.Query(handler.Action)
	switch action {
	// 预览技术资料
	case handler.ViewData:
		handle.GetData(ctx, id)
	// 查看资料简介
	case handler.Introduction:
		intro, err := handle.GetIntroduction(id)
		if err != nil {
			response.FailedResponse(ctx, err)
			return
		}
		response.SuccessResponse(ctx, intro)
	// 下载技术资料
	case handler.DownLoad:
		res, _, err := dal.QueryDataByCondition(&request.QueryDataForRootReq{
			ID:         id,
			Limit:      1,
			DataType:   utils.DataTypeAll,
			Department: utils.DepartmentAll,
			AuthLevel:  utils.AuthLevelAll,
		})
		if err != nil {
			response.FailedResponse(ctx, err)
			return
		}
		if len(res) != 1 {
			response.FailedResponse(ctx, fmt.Errorf("len(res) is not only"))
			return
		}
		ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", res[0].Title)) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
		ctx.File(res[0].DataUrl)
		return
	default:
		response.FailedResponse(ctx, fmt.Errorf("action: %s 暂不支持", action))
		return
	}
}

func OperateTechniqueData(ctx *gin.Context) {
	// 获取action和请求体
	action := ctx.Query(handler.Action)
	idStr := ctx.Param(handler.ParamID)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("atoi idStr failed: ", err.Error())
		response.FailedResponse(ctx, err)
		return
	}
	switch action {
	case handler.DeleteData:
		err = handle.DoDeleteData(id)
		if err != nil {
			log.Println("DoDeleteData failed: ", err.Error())
			response.FailedResponse(ctx, err)
			return
		}
		response.SuccessResponse(ctx, nil)
		return
	// 更新技术资料
	case handler.UpdateData:
		userName, _ := ctx.Get("user_name")
		content, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			response.FailedResponse(ctx, err)
			return
		}
		updateParam := &request.UpdateDataForRootReq{}
		err = json.Unmarshal(content, updateParam)
		if err != nil {
			response.FailedResponse(ctx, err)
			return
		}
		file, err := ctx.FormFile(updateParam.Title)
		if err != nil {
			response.FailedResponse(ctx, err)
			return
		}
		err = handle.DoUpdateData(userName.(string), id, updateParam, file)
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
