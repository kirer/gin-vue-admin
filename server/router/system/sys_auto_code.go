package system

import (
	"github.com/gin-gonic/gin"
	v1 "kirer.cn/server/api/v1"
)

type AutoCodeRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup) {
	router := Router.Group("autoCode")
	api := v1.ApiGroupApp.SystemApiGroup.AutoCodeApi
	{
		router.GET("getDB", api.GetDB)                  // 获取数据库
		router.GET("getTables", api.GetTables)          // 获取对应数据库的表
		router.GET("getColumn", api.GetColumn)          // 获取指定表所有字段信息
		router.POST("preview", api.PreviewTemp)         // 获取自动创建代码预览
		router.POST("createTemp", api.CreateTemp)       // 创建自动化代码
		router.POST("createPackage", api.CreatePackage) // 创建package包
		router.POST("getPackage", api.GetPackage)       // 获取package包
		router.POST("delPackage", api.DelPackage)       // 删除package包
		router.POST("createPlug", api.AutoPlug)         // 自动插件包模板
		router.POST("installPlugin", api.InstallPlugin) // 自动安装插件
		router.POST("pubPlug", api.PubPlug)             // 打包插件
	}
}
