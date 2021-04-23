package main

import (
	"course/config"
	"course/dal"
	"course/data"
	"course/handler"
	"course/handler/normal"
	"course/handler/root"
	"course/logs"
	"course/response"
	"course/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	_, ok := os.LookupEnv("CUR_ENV")
	if ok {
		log.Println("no CUR_ENV")
	}
	logs.InitLog()             // 初始化日志
	log.Println("Init config") // 读取配置内容
	totalConf := &config.TotalConf{}
	err := utils.UnmarshalConf(totalConf)
	if err != nil {
		log.Println("UnmarshalConf failed: ", err.Error())
		return
	}
	log.Println("Init MySQL")
	err = dal.InitMySQL(totalConf.MySQL) // 初始化数据库
	if err != nil {
		log.Println("Init MySQL failed: ", err.Error())
		return
	}
	log.Println("Init dataPath")

	err = data.InitDataDir(totalConf.DataPath) // 初始化数据目录
	if err != nil {
		log.Println("Init DataDir failed: ", err.Error())
		return
	}

	r := gin.Default()                           // 获取engine，启动路由
	r.Use(getSessionMidWare(), getAuthMidWare()) // 获取Session中间件和自定义中间件
	log.Println("Init routers")
	initAllRouters(r)
	err = r.Run(":80")
	if err != nil {
		log.Println("gin engine run failed: ", err.Error())
		return
	}
}

func getSessionMidWare() gin.HandlerFunc {
	store := cookie.NewStore([]byte("course"))
	midWare := sessions.Sessions("mySession", store)
	return midWare
}

// 自定义中间件，拦截http请求，校验是否处于登录状态下
func getAuthMidWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 如果是登陆路由，直接返回
		if ctx.FullPath() == "/common/login" || ctx.FullPath() == "/addroot" {
			return
		}
		session := sessions.Default(ctx)
		val := session.Get("is_login")
		// 如果没有查询到user_name，直接返回
		if val == nil {
			ctx.Abort()
			response.FailedResponse(ctx, fmt.Errorf("please login first"))
			return
		}
		// 如果类型转换失败，直接返回
		userNameStr, flag := val.(string)
		if !flag {
			ctx.Abort()
			response.FailedResponse(ctx, fmt.Errorf("please login first"))
			return
		}
		// 如果session中获取到的user_name和请求参数中的user_name不一致，直接返回
		userName := ctx.Query("user_name")
		if userName != userNameStr {
			ctx.Abort()
			response.FailedResponse(ctx, fmt.Errorf("count auth failed, please login again"))
			return
		}
		// 登录验证成功
		ctx.Set("user_name", userName)
	}
}

func initAllRouters(r *gin.Engine) {
	r.GET("/statusinfo", handler.GetStatusInfo) // 获取全局变量，映射关系，前端使用 -- done
	r.POST("/addroot", handler.AddRootCount)    // 添加管理员账户 -- done
	rootGroup := r.Group("/root")               // 管理员路由组 -- done
	userGroup := r.Group("/user")               // 普通用户路由组 -- done
	commonGroup := r.Group("/common")           // 通用路由组 -- 主要是资料预览，下载，简介查看这类root和user通用的请求
	initRootGroup(rootGroup)
	initUserGroup(userGroup)
	initCommonGroup(commonGroup)
}

func initRootGroup(rootGroup *gin.RouterGroup) {
	rootGroup.POST("/queryusers", root.QueryUserCount)     // 查询用户账户 -- done -- tested
	rootGroup.POST("/addusers", root.AddUserCount)         // 增加用户账户 -- done -- tested
	rootGroup.POST("/user/:id", root.OperateUserCount)     // 对单个用户账户进行操作 -- done -- tested
	rootGroup.POST("/querydatas", root.QueryTechniqueData) // 查询技术资料 -- done -- tested
	rootGroup.POST("/adddatas", root.AddTechniqueData)     // 增加技术资料 -- done -- tested
	rootGroup.POST("/data/:id", root.OperateTechniqueData) // 对技术资料进行更新/删除操作 -- done
}

func initUserGroup(userGroup *gin.RouterGroup) {
	userGroup.POST("/datas", normal.QueryTechniqueData)      // 检索技术资料 -- done -- tested
	userGroup.POST("/data/:id", normal.GetTechniqueDataInfo) // 查看技术资料 -- done -- tested
	userGroup.POST("/changepw", normal.ChangePassWord)       // 修改密码 -- done -- tested
}

// 预览技术资料，查看资料简介，下载技术资料
func initCommonGroup(commonGroup *gin.RouterGroup) {
	commonGroup.POST("/data/:id", root.GetTechniqueDataInfo) // 预览技术资料和资料简介/下载技术资料 -- done -- tested
	commonGroup.POST("/logout", handler.LogOut)              // 退出登录
	commonGroup.POST("/login", handler.Login)                // 登录 -- done -- tested
}
