package handle

import (
	"course/dal"
	"course/data"
	"course/request"
	"course/response"
	"course/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"mime/multipart"
)

func DoQueryData(queryParam *request.QueryDataForRootReq) (*response.QueryDataForRootResp, error) {
	datas, count, err := dal.QueryDataByCondition(queryParam)
	if err != nil {
		log.Println("QueryDataByCondition failed: ", err.Error())
		return nil, err
	}
	resp := &response.QueryDataForRootResp{
		Count: count,
	}
	res := make([]*response.Data, 0)
	for i := 0; i < len(datas); i++ {
		res = append(res, transDBDataToFeDataRoot(datas[i]))
	}
	resp.Res = res
	return resp, nil
}

// FBI Warning: 应有但无(应该有一个事务，但是没必要加)
func DoAddData(userName string, addParam *request.AddDataForRootReq, fileHeader *multipart.FileHeader) error {
	if err := checkAddDataParam(addParam); err != nil {
		return err
	}
	url, err := data.Save(fileHeader)
	if err != nil {
		return err
	}
	addData := &dal.TechniqueData{
		AddUser:      userName,
		UpdateUser:   userName,
		Type:         addParam.DataType,
		AuthLevel:    addParam.AuthLevel,
		Department:   addParam.Department,
		Introduction: addParam.Introduction,
		DataUrl:      url,
		Title:        addParam.Title,
	}
	return dal.AddData(addData)
}

// FBI Warning: 讲道理需要用一个事务来保证原子性，那什么叫毕设啊？完全没必要！！！！
func DoUpdateData(userName string, updateID int, updateParam *request.UpdateDataForRootReq, fileHeader *multipart.FileHeader) error {
	// 存储新的文件
	url, err := data.Save(fileHeader)
	if err != nil {
		log.Println("save file failed: ", err.Error())
		return err
	}
	res, _, err := dal.QueryDataByCondition(&request.QueryDataForRootReq{
		ID: updateID,
	})
	if err != nil {
		log.Println("query data[ ", updateID, " ] failed: ", err.Error())
		return err
	}
	if len(res) != 1 {
		return fmt.Errorf("data[ %s] record is required one", updateID)
	}
	// 删除老文件
	err = data.Remove(res[0].DataUrl)
	if err != nil {
		return err
	}
	// 更新数据库
	return dal.UpdateDataByDataID(updateID, updateParam, url, userName)
}

// FBI Warning: 按道理来说这里应该开一个事务保证两者的原子性
// 但是，毕设有必要么？完全没有！！！
func DoDeleteData(dataID int) error {
	// 删除磁盘文件
	res, _, err := dal.QueryDataByCondition(&request.QueryDataForRootReq{ID: dataID})
	if err != nil {
		log.Println("QueryDataByCondition failed: ", err.Error())
		return err
	}
	if len(res) != 1 {
		log.Println("QueryDataByCondition failed: more than one records")
		return fmt.Errorf("id[ %s ] records are more then one", dataID)
	}
	record := res[0]
	err = data.Remove(record.DataUrl)
	if err != nil {
		log.Println("Remove id[ ", dataID, " ] filed: ", err.Error())
		return err
	}
	// 删除数据库记录
	err = dal.DeleteDataByID(dataID)
	if err != nil {
		log.Println("DeleteDataByID failed: ", err.Error())
		return err
	}
	return nil
}

func GetData(ctx *gin.Context, id int) {
	// resp := &response.ViewDataResp{}
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
		response.FailedResponse(ctx, fmt.Errorf("res is not only"))
		return
	}
	bytes, err := ioutil.ReadFile(res[0].DataUrl)
	if err != nil {
		response.FailedResponse(ctx, err)
		return
	}
	if res[0].Type == int(utils.TypeVideo) {
		ctx.Writer.Header().Add("Content-Type", "video/mp4")
	} else if res[0].Type == int(utils.TypePdf) {
		ctx.Writer.Header().Add("Content-Type", "application/pdf")
	} else {
		ctx.Writer.Header().Add("Content-Type", "application/msword")
	}
	_, _ = ctx.Writer.Write(bytes)
}

func GetIntroduction(id int) (string, error) {
	res, _, err := dal.QueryDataByCondition(&request.QueryDataForRootReq{
		ID:         id,
		Limit:      1,
		DataType:   utils.DataTypeAll,
		Department: utils.DepartmentAll,
		AuthLevel:  utils.AuthLevelAll,
	})
	if err != nil {
		return "", err
	}
	if len(res) != 1 {
		return "", fmt.Errorf("res is not only")
	}
	return res[0].Introduction, nil
}

func checkAddDataParam(addParam *request.AddDataForRootReq) error {
	if addParam == nil {
		return fmt.Errorf("a nil param is not allowed")
	}
	if addParam.Title == "" {
		return fmt.Errorf("Title Field is not allowed nil")
	}
	if addParam.AuthLevel < 0 {
		return fmt.Errorf("AuthLevel Field is required to be 0")
	}
	if addParam.Introduction == "" {
		return fmt.Errorf("Introduction is not allow nil")
	}
	typeCode := utils.DataTypeCode(addParam.DataType)
	if typeCode != utils.TypePdf && typeCode != utils.TypeWord && typeCode != utils.TypeVideo {
		return fmt.Errorf("Data Type is only allowed in [%s, %s, %s]", utils.TypePdf, utils.TypeWord, utils.TypeVideo)
	}
	return nil
}

func transDBDataToFeDataRoot(data *dal.TechniqueData) *response.Data {
	return &response.Data{
		ID:               int(data.ID),
		Title:            data.Title,
		CreateTime:       data.CreatedAt,
		AddUser:          data.AddUser,
		LastModifiedTime: data.UpdatedAt,
		LastModifiedUser: data.UpdateUser,
		DataType:         utils.DataTypeMap[utils.DataTypeCode(data.Type)],
		TypeCode:         data.Type,
		AuthLevel:        data.AuthLevel,
		Department:       utils.DepartmentMap[utils.DepartmentCode(data.Department)],
		DepartmentCode:   data.Department,
	}
}
