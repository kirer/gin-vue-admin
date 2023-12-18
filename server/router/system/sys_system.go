package system

import (
	"github.com/gin-gonic/gin"
	v1 "kirer.cn/server/api/v1"
	"kirer.cn/server/middleware"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	router := Router.Group("system").Use(middleware.Record())
	api := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		router.POST("get_config", api.GetConfig) // 获取配置文件内容
		router.POST("set_config", api.SetConfig) // 设置配置文件内容
		router.POST("get_info", api.GetInfo)     // 获取服务器信息
		router.POST("reload", api.Reload)        // 重启服务
	}
}
