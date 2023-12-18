package system

import (
	"github.com/gin-gonic/gin"
	v1 "kirer.cn/server/api/v1"
	"kirer.cn/server/middleware"
)

type DicRouter struct{}

func (s *DicRouter) InitSysDicRouter(Router *gin.RouterGroup) {
	router := Router.Group("dic").Use(middleware.Record())
	router2 := Router.Group("dic")
	api := v1.ApiGroupApp.SystemApiGroup.DicApi
	{
		router.POST("create", api.Create)   // 新建SysDic
		router.DELETE("delete", api.Delete) // 删除SysDic
		router.PUT("update", api.Update)    // 更新SysDic
	}
	{
		router2.GET("get", api.Get)   // 根据ID获取SysDic
		router2.GET("list", api.List) // 获取SysDic列表
	}
}
