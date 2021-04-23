package request

// 管理员账户查询用户账户请求参数
type QueryUserCountReq struct {
	UserName   string `json:"user_name"`  // 用户名
	Department int    `json:"department"` // 隶属部门
	Status     int    `json:"status"`     // 帐号状态
	UserType   int    `json:"user_type"`  // 员工类型
	AuthLevel  int    `json:"auth_level"` // 权限等级
	Offset     int    `json:"offset"`     // 偏移量 - 分页
	Limit      int    `json:"limit"`      // 每一页个数 - 分页
}

// 管理员账户增加用户账户请求参数
type AddUserCountReq struct {
	UserName   string `json:"user_name"`  // 用户名称
	PassWord   string `json:"pass_word"`  // 密码
	UserType   int    `json:"user_type"`  // 员工类型
	AuthLevel  int    `json:"auth_level"` // 权限等级
	Department int    `json:"department"` // 隶属部门
	Status     int    `json:"status"`     // 帐号状态
}

// 管理员账户修改用户账户请求参数
type UpdateUserCountReq struct {
	AddUserCountReq
}
