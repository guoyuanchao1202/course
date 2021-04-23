package request

// 普通用户查询技术资料请求参数
type QueryDataForUserReq struct {
	QueryStr  string `json:"query_str"`  // 查询内容
	QueryType int    `json:"query_type"` // 查询方式
	Offset    int    `json:"offset"`     // 偏移量 - 分页
	Limit     int    `json:"limit"`      // 每一页个数 - 分页
}

// 普通用户账户修改密码
type ChangePassWordForUser struct {
	OldPassWord string `json:"old_pass_word"` // 旧密码
	NewPassWord string `json:"new_pass_word"` // 新密码
}

type LoginReq struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	IsRoot   int    `json:"is_root"`
}
