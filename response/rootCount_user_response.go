// 管理员账户操作用户账户相应参数
package response

// 管理员账户查询用户账户响应参数
type User struct {
	ID             int    `json:"id"`              // 用户唯一ID
	UserName       string `json:"user_name"`       // 用户名
	PassWord       string `json:"pass_word"`       // 用户密码
	UserType       string `json:"user_type"`       // 员工类型(技术 / 设计..)
	TypeCode       int    `json:"type_code"`       // 员工类型码
	AuthLevel      int    `json:"auth_level"`      // 员工权限等级
	Status         string `json:"status"`          // 账号状态(已禁用 / 正常状态)
	StatusCode     int    `json:"status_code"`     // 帐号状态码
	Department     string `json:"department"`      // 员工隶属部门
	DepartmentCode int    `json:"department_code"` // 部门码
}

type QueryUserRespData struct {
	Count int     `json:"count"`
	Res   []*User `json:"res"`
}
