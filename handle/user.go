package handle

import (
	"course/dal"
	"course/request"
	"course/response"
	"course/utils"
	"fmt"
	"log"
)

func DoChangePW(userName string, changeParam *request.ChangePassWordForUser) error {
	res, _, err := dal.QueryUserCountByCondition(&request.QueryUserCountReq{
		UserName: userName,
	})
	if err != nil {
		return err
	}
	if len(res) != 1 {
		return fmt.Errorf("len(res) != 1")
	}
	if res[0].PassWord != changeParam.OldPassWord {
		return fmt.Errorf("oldPassword is incorrect")
	}
	return dal.UpdateUserPw(userName, changeParam.NewPassWord)
}

//AddUser    string `json:"add_user"`   // 添加人
//	DataType   int    `json:"data_type"`  // 数据类型
//	Department int    `json:"department"` // 隶属部门
//	AuthLevel  int    `json:"auth_level"` // 权限等级
func DoQueryDataByID(queryID int) (*response.QueryDataForUserResp, error) {
	res, _, err := dal.QueryDataByCondition(&request.QueryDataForRootReq{
		ID:         queryID,
		Limit:      1,
		DataType:   utils.DataTypeAll,
		Department: utils.DepartmentAll,
		AuthLevel:  utils.AuthLevelAll,
	})
	if err != nil {
		log.Println("QueryDataByCondition failed: ", err.Error())
		return nil, err
	}
	if len(res) != 1 {
		log.Println("QueryDataByCondition failed: len(res) = ", len(res))
		return nil, fmt.Errorf("len(res) is %s, required 1", len(res))
	}
	resp := transDBToFe(res[0])
	return resp, nil
}

func DoQueryDataByCond(userName string, queryParam *request.QueryDataForUserReq) (*response.QueryDatasForUser, error) {
	user, err := dal.QueryUserByUserName(userName)
	if err != nil {
		log.Println("QueryUserByUserName failed: ", err.Error())
		return nil, err
	}
	queryRes, count, err := dal.QueryDataWithFullText(user.AuthLevel, queryParam)
	if err != nil {
		log.Println("QueryDataWithFullText failed: ", err.Error())
		return nil, err
	}
	resp := &response.QueryDatasForUser{Count: count}
	res := make([]*response.QueryDataForUserResp, 0)
	for i := 0; i < count; i++ {
		res = append(res, transDBToFe(queryRes[i]))
	}
	resp.Res = res
	return resp, nil
}

func transDBToFe(data *dal.TechniqueData) *response.QueryDataForUserResp {
	return &response.QueryDataForUserResp{
		ID:           int(data.ID),
		Title:        data.Title,
		Introduction: data.Introduction,
		DataType:     utils.DataTypeMap[utils.DataTypeCode(data.Type)],
		TypeCode:     data.Type,
		Department:   utils.DepartmentMap[utils.DepartmentCode(data.Department)],
		AddUser:      data.AddUser,
		AddTime:      data.CreatedAt,
	}
}
