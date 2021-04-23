package request

// 管理员账户查询技术资料请求参数
type QueryDataForRootReq struct {
	ID         int    `json:"id"`         // 资料ID
	AddUser    string `json:"add_user"`   // 添加人
	DataType   int    `json:"data_type"`  // 数据类型
	Department int    `json:"department"` // 隶属部门
	AuthLevel  int    `json:"auth_level"` // 权限等级
	Offset     int    `json:"offset"`     // 偏移量 - 分页
	Limit      int    `json:"limit"`      // 每一页个数 - 分页
}

// 管理员账户增加技术资料请求参数
type AddDataForRootReq struct {
	Title        string `json:"title"`        // 资料标题
	DataType     int    `json:"data_type"`    // 资料类型
	AuthLevel    int    `json:"auth_level"`   // 权限等级
	Department   int    `json:"department"`   // 隶属部门
	Introduction string `json:"introduction"` // 资料概述
}

// 管理员账户修改技术资料请求参数
type UpdateDataForRootReq struct {
	AddDataForRootReq
}
