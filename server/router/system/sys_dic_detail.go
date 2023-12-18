package system

import (
	"github.com/gin-gonic/gin"
	v1 "kirer.cn/server/api/v1"
	"kirer.cn/server/middleware"
)

type DicDetailRouter struct{}

func (s *DicDetailRouter) InitSysDicDetailRouter(Router *gin.RouterGroup) {
	router := Router.Group("dic_detail").Use(middleware.Record())
	router2 := Router.Group("dic_detail")
	api := v1.ApiGroupApp.SystemApiGroup.DicDetailApi
	{
		router.POST("create", api.Create)   // 新建
		router.DELETE("delete", api.Delete) // 删除
		router.PUT("update", api.Update)    // 更新
	}
	{
		router2.GET("get", api.Get)   // 根据ID获取
		router2.GET("list", api.List) // 获取列表
	}
}
