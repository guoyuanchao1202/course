package handle

import (
	"course/dal"
	"course/request"
	"course/response"
	"course/utils"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

func DoQueryUser(req *request.QueryUserCountReq) (*response.QueryUserRespData, error) {
	users, count, err := dal.QueryUserCountByCondition(req)
	if err != nil {
		log.Println("query user counts failed: ", err.Error())
		return nil, err
	}
	resp := &response.QueryUserRespData{
		Count: count,
	}
	res := make([]*response.User, 0)
	for i := 0; i < len(users); i++ {
		res = append(res, transDBUserToFeUser(users[i]))
	}
	resp.Res = res
	return resp, nil
}

func DoAddUser(addParam *request.AddUserCountReq) error {
	if err := checkAddParam(addParam); err != nil {
		return err
	}
	user := &dal.User{
		Model:      gorm.Model{},
		UserName:   addParam.UserName,
		PassWord:   addParam.PassWord,
		EmpType:    addParam.UserType,
		AuthLevel:  addParam.AuthLevel,
		Status:     addParam.Status,
		Department: addParam.Department,
	}
	return dal.AddUserCount(user)
}

func DoUpdateUser(id int, content []byte) error {
	updateParam := &request.UpdateUserCountReq{}
	err := json.Unmarshal(content, updateParam)
	if err != nil {
		log.Println("Unmarshal updateParam failed: ", err.Error())
		return err
	}
	return dal.UpdateUserCountByID(id, updateParam)
}

func DoDeleteUser(id int) error {
	return dal.DeleteUserCountByID(id)
}

func DoChangeStatus(id int, content []byte) error {
	userStatus := &struct {
		Enabled int `json:"enabled"`
	}{}
	err := json.Unmarshal(content, userStatus)
	if err != nil {
		log.Println("Unmarshal changeStatus failed: ", err.Error())
		return err
	}
	return dal.ChangeUserCountStatusByID(id, userStatus.Enabled)
}

func checkAddParam(addParam *request.AddUserCountReq) error {
	if addParam.UserName == "" || addParam.PassWord == "" {
		return fmt.Errorf("userName | passWord is not allowed nil")
	}
	if utils.EmpTypeMap[utils.UserTypeCode(addParam.UserType)] == "" {
		return fmt.Errorf("invalid empTypeCode: %s", addParam.UserType)
	}
	if utils.DepartmentMap[utils.DepartmentCode(addParam.Department)] == "" {
		return fmt.Errorf("invalid departmentCode: %s", addParam.Department)
	}
	if addParam.Status != utils.Enabled && addParam.Status != utils.Disabled {
		return fmt.Errorf("invalid statusCode: %s. only allowed [%s | %s]", addParam.Status, utils.Enabled, utils.Disabled)
	}
	return nil
}

func transDBUserToFeUser(user *dal.User) *response.User {
	return &response.User{
		ID:             int(user.ID),
		UserName:       user.UserName,
		PassWord:       user.PassWord,
		UserType:       utils.EmpTypeMap[utils.UserTypeCode(user.EmpType)],
		TypeCode:       user.EmpType,
		AuthLevel:      user.AuthLevel,
		Status:         utils.CountStatusMap[user.Status],
		StatusCode:     user.Status,
		Department:     utils.DepartmentMap[utils.DepartmentCode(user.Department)],
		DepartmentCode: user.Department,
	}
}
