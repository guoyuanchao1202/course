package dal

import (
	"course/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var Dao *gorm.DB

// 初始化MySQL
func InitMySQL(mysqlConf config.MySQLConfig) (err error) {
	format := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
	log.Println(format)
	// 用户名-密码-主机名-端口号-数据库名
	DSN := fmt.Sprintf(format, mysqlConf.DBUser, mysqlConf.DBPassWord, mysqlConf.DBHost, mysqlConf.DBPort, mysqlConf.DBName)
	Dao, err = gorm.Open("mysql", DSN)
	if err != nil {
		return
	}
	Dao = Dao.LogMode(true)
	return
}
