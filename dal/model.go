package dal

import "github.com/jinzhu/gorm"

const (
	id                 = "id"
	createdAt          = "created_at"
	updatedAt          = "updated_at"
	deletedAt          = "deleted_at"
	userName           = "user_name"
	passWord           = "pass_word"
	empType            = "emp_type"
	authLevel          = "auth_level"
	status             = "status"
	department         = "department"
	addUser            = "add_user"
	updateUser         = "update_user"
	Type               = "type"
	introduction       = "introduction"
	dataUrl            = "data_url"
	tableUser          = "users"
	isRoot             = "is_root"
	tableTechniqueData = "technique_datas"
	title              = "title"
)

type User struct {
	gorm.Model
	UserName   string
	PassWord   string
	EmpType    int
	AuthLevel  int
	Status     int
	Department int
	IsRoot     int
}
func (user User)TableName() string {
	return tableUser
}


type TechniqueData struct {
	gorm.Model
	AddUser      string
	UpdateUser   string
	Type         int
	AuthLevel    int
	Department   int
	Introduction string
	DataUrl      string
	Title        string
}

func (data TechniqueData)TableName() string {
	return tableTechniqueData
}
