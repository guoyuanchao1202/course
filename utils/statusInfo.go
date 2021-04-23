package utils

const (
	Enabled    = 1
	Disabled   = 0
	FeEnabled  = "正常"
	FeDisabled = "禁用"

	StatusAll = -1
)

// 账号状态映射
var CountStatusMap = map[int]string{
	Enabled:  FeEnabled,
	Disabled: FeDisabled,
}

type UserTypeCode int

const (
	TypeTechnician          UserTypeCode = iota // 技术员工
	TypeProductStaff                            // 产品员工
	TypeDesignStaff                             // 设计员工
	TypePlanningStaff                           // 策划员工
	TypeAdministrativeStaff                     // 行政员工

	FeTypeTechnician      = "技术序列"
	FeTypeProductStaff    = "产品序列"
	FeTypeDesignStaff     = "设计序列"
	FeTypePlanningStaff   = "策划序列"
	FeAdministrativeStaff = "行政序列"

	TypeOfAllUser = -1
)

// 员工类型映射关系
var EmpTypeMap = map[UserTypeCode]string{
	TypeTechnician:          FeTypeTechnician,
	TypeProductStaff:        FeTypeProductStaff,
	TypeDesignStaff:         FeTypeDesignStaff,
	TypePlanningStaff:       FeTypePlanningStaff,
	TypeAdministrativeStaff: FeAdministrativeStaff,
}

type DataTypeCode int

const (
	TypePdf   DataTypeCode = iota // pdf资料
	TypeWord                      // 文档资料
	TypeVideo                     // 视频资料

	FeTypePdf   = "pdf"
	FeTypeWord  = "word文档"
	FeTypeVideo = "视频"

	DataTypeAll = -1
)

// 资料映射关系
var DataTypeMap = map[DataTypeCode]string{
	TypePdf:   FeTypePdf,
	TypeWord:  FeTypeWord,
	TypeVideo: FeTypeVideo,
}

type DepartmentCode int

const (
	ProductDepartment        DepartmentCode = iota // 产品部
	DesignDepartment                               // 设计部
	AdministrationDepartment                       // 行政部
	OperationDepartment                            // 运营部

	FeProductDepartment        = "产品部"
	FeDesignDepartment         = "设计部"
	FeAdministrationDepartment = "行政部"
	FeOperationDepartment      = "运营部"

	DepartmentAll = -1
)

// 部门映射关系
var DepartmentMap = map[DepartmentCode]string{
	ProductDepartment:        FeProductDepartment,
	DesignDepartment:         FeDesignDepartment,
	AdministrationDepartment: FeAdministrationDepartment,
	OperationDepartment:      FeOperationDepartment,
}

const AuthLevelAll  = -1