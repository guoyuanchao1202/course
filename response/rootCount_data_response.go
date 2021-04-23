// 管理员账户操作技术资料相应参数
package response

import "time"

// 管理员账户查询资料响应参数
type Data struct {
	ID               int       `json:"id"`                 // 资料唯一ID
	Title            string    `json:"title"`              // 资料名称
	CreateTime       time.Time `json:"create_time"`        // 资料创建时间
	AddUser          string    `json:"add_user"`           // 资料添加人
	LastModifiedTime time.Time `json:"last_modified_time"` // 资料最后修改时间
	LastModifiedUser string    `json:"last_modified_user"` // 资料最后修改人
	DataType         string    `json:"data_type"`          // 资料类型(pdf / 视频 / 文档 ...)
	TypeCode         int       `json:"type_code"`          // 资料类型码
	AuthLevel        int       `json:"auth_level"`         // 资料权限等级
	Department       string    `json:"department"`         // 资料隶属部门(产品部 / 设计部 ...)
	DepartmentCode   int       `json:"department_code"`    // 资料隶属部门代码
}

type QueryDataForRootResp struct {
	Count int     `json:"count"`
	Res   []*Data `json:"res"`
}

// 预览资料响应参数
type ViewDataResp struct {
	DataType int    // 资料类型
	Data     []byte // 资料内容
}
