package response

import "course/utils"

type StatusInfoResp struct {
	CountStatusMap map[int]string                  `json:"count_status_map"`
	EmpTypeMap     map[utils.UserTypeCode]string   `json:"emp_type_map"`
	DataTypeMap    map[utils.DataTypeCode]string   `json:"data_type_map"`
	DepartmentMap  map[utils.DepartmentCode]string `json:"department_map"`
}
