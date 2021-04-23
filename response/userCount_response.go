// 用户账户操作资料相应参数
package response

import "time"

// 用户账户查询技术资料响应参数
type QueryDataForUserResp struct {
	ID           int       `json:"id"`           // 资料唯一ID
	Title        string    `json:"title"`        // 资料名称
	Introduction string    `json:"introduction"` // 资料简介
	DataType     string    `json:"data_type"`    // 资料类型
	Department   string    `json:"department"`   // 资料所属部门
	TypeCode     int       `json:"type_code"`    // 资料类型码
	AddUser      string    `json:"add_user"`     // 资料添加人
	AddTime      time.Time `json:"add_time"`     // 资料添加时间
}

type QueryDatasForUser struct {
	Res   []*QueryDataForUserResp `json:"res"`
	Count int                     `json:"count"`
}

type LoginResp struct {
	UserName string `json:"user_name"`
	IsRoot   int    `json:"is_root"`
}
