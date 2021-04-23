package dal

import (
	"course/request"
	"course/utils"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/didi/gendry/builder"
	"github.com/jinzhu/gorm"
	"log"
)

func QueryUserByUserName(loginUserName string) (*User, error) {
	resUser := &User{}
	err := Dao.Where(userName+" = ?", loginUserName).First(resUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return resUser, err
}

func QueryUserCountByCondition(query *request.QueryUserCountReq) ([]*User, int, error) {
	res := make([]*User, 0)
	// 获取where条件
	where := getUserCountQueryMap(query)
	// 构建sql语句
	buildSql, args, err := where.ToSql()
	if err != nil {
		return nil, 0, err
	}
	err = Dao.Where(buildSql, args...).Limit(query.Limit).Offset(query.Offset).Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("query user failed: ", err.Error())
		return nil, -1, err
	}
	if err == gorm.ErrRecordNotFound {
		return res, 0, nil
	}
	var count int
	err = Dao.Model(&User{}).Where(buildSql, args...).Count(&count).Error
	return res, count, err
}

func QueryDataByCondition(query *request.QueryDataForRootReq) ([]*TechniqueData, int, error) {
	res := make([]*TechniqueData, 0)
	where := getRootDataQueryMap(query)
	buildSql, args, err := where.ToSql()
	if err != nil {
		return nil, 0, err
	}
	err = Dao.Where(buildSql, args...).Limit(query.Limit).Offset(query.Offset).Find(&res).Error
	if err != nil {
		log.Println("root query datas failed: ", err.Error())
		return nil, -1, err
	}
	var count int
	err = Dao.Model(&TechniqueData{}).Where(buildSql, args...).Count(&count).Error
	return res, count, err
}

//	QueryStr  string `json:"query_str"`  // 查询内容
//	QueryType int    `json:"query_type"` // 查询方式
//	Offset    int    `json:"offset"`     // 偏移量 - 分页
//	Limit     int    `json:"limit"`      // 每一页个数 - 分页
//  MATCH(`name`,`address`) AGAINST('聪 广东')
func QueryDataWithFullText(QueryAuthLevel int, query *request.QueryDataForUserReq) ([]*TechniqueData, int, error) {
	res := make([]*TechniqueData, 0)
	var count int
	where := fmt.Sprintf("MATCH(%s, %s) AGAINST (\"%s\") AND %s >= %d", title, introduction, query.QueryStr, authLevel, QueryAuthLevel)
	var err error
	if query.QueryType != -1 {
		where = fmt.Sprintf("%s AND %s = %d", where, Type, query.QueryType)
	}
	err = Dao.Where(where).Limit(query.Limit).Offset(query.Offset).Find(&res).Error
	if err != nil {
		return nil, 0, err
	}
	err = Dao.Model(&TechniqueData{}).Where(where).Limit(query.Limit).Offset(query.Offset).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

func AddUserCount(user *User) error {
	return Dao.Create(user).Error
}

func AddData(data *TechniqueData) error {
	return Dao.Create(data).Error
}

func DeleteDataByID(deleteID int) error {
	return Dao.Where(id+" = ?", deleteID).Delete(&TechniqueData{}).Error
}

func UpdateUserCountByID(userID int, update *request.UpdateUserCountReq) error {
	where := map[string]interface{}{
		id: userID,
	}
	updateCond := map[string]interface{}{
		userName:   update.UserName,
		passWord:   update.PassWord,
		empType:    update.UserType,
		authLevel:  update.AuthLevel,
		status:     update.Status,
		department: update.Department,
	}
	cond, vals, err := builder.BuildUpdate(tableUser, where, updateCond)
	if err != nil {
		log.Print("buildUpdate failed: ", err.Error())
		return err
	}
	return Dao.Exec(cond, vals...).Error
}

func DeleteUserCountByID(userID int) error {
	cond, vals, err := builder.BuildDelete(tableUser, map[string]interface{}{id: userID})
	if err != nil {
		log.Println("buildDelete failed: ", err)
		return err
	}
	return Dao.Exec(cond, vals...).Error
}

func ChangeUserCountStatusByID(userID int, newStatus int) error {
	where := map[string]interface{}{
		id: userID,
	}
	updateCond := map[string]interface{}{
		status: newStatus,
	}
	cond, vals, err := builder.BuildUpdate(tableUser, where, updateCond)
	if err != nil {
		log.Println("buildUpdate failed: ", err.Error())
		return err
	}
	return Dao.Exec(cond, vals...).Error
}
func UpdateDataByDataID(updateID int, updateParam *request.UpdateDataForRootReq, url, updateUserName string) error {
	where := map[string]interface{}{
		id: updateID,
	}
	updateCond := map[string]interface{}{
		updateUser:   updateUserName,
		dataUrl:      url,
		title:        updateParam.Title,
		Type:         updateParam.DataType,
		authLevel:    updateParam.AuthLevel,
		department:   updateParam.Department,
		introduction: updateParam.Introduction,
	}
	cond, vals, err := builder.BuildUpdate(tableTechniqueData, where, updateCond)
	if err != nil {
		log.Println("buildUpdate failed: ", err.Error())
		return err
	}
	return Dao.Exec(cond, vals...).Error
}

func UpdateUserPw(changeUserName, passWd string) error {
	where := map[string]interface{}{
		userName: changeUserName,
	}
	updateCond := map[string]interface{}{
		passWord: passWd,
	}
	cond, vals, err := builder.BuildUpdate(tableUser, where, updateCond)
	if err != nil {
		return err
	}
	return Dao.Exec(cond, vals...).Error
}

func getUserCountQueryMap(query *request.QueryUserCountReq) squirrel.Eq {
	// where := make(map[string]interface{})
	where := squirrel.Eq{}
	if query.UserName != "" {
		where[userName] = query.UserName
	}
	if query.Department != utils.DepartmentAll {
		where[department] = query.Department
	}
	if query.Status != utils.StatusAll {
		where[status] = query.Status
	}
	if query.UserType != utils.TypeOfAllUser {
		where[empType] = query.UserType
	}
	if query.AuthLevel != utils.AuthLevelAll {
		where[authLevel] = query.AuthLevel
	}
	//if query.Limit != 0 {
	//	where["_limit"] = []uint{uint(query.Offset), uint(query.Limit)}
	//}
	return where
}

func getRootDataQueryMap(query *request.QueryDataForRootReq) squirrel.Eq {
	where := squirrel.Eq{}
	if query.ID != 0 {
		where[id] = query.ID
	}
	if query.AddUser != "" {
		where[addUser] = query.AddUser
	}
	if query.DataType != utils.DataTypeAll {
		where[Type] = query.DataType
	}
	if query.Department != utils.DepartmentAll {
		where[department] = query.Department
	}
	if query.AuthLevel != utils.AuthLevelAll {
		where[authLevel] = query.AuthLevel
	}
	return where
}
