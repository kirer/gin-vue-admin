package system

import (
	"github.com/gin-gonic/gin"
	v1 "kirer.cn/server/api/v1"
)

type RecordRouter struct{}

func (s *RecordRouter) InitSysRecordRouter(Router *gin.RouterGroup) {
	router := Router.Group("record")
	api := v1.ApiGroupApp.SystemApiGroup.RecordApi
	{
		router.POST("create", api.Create)     // 新建
		router.DELETE("delete", api.Delete)   // 删除
		router.DELETE("deletes", api.Deletes) // 批量删除
		router.GET("get", api.Get)            // 根据ID获取
		router.GET("list", api.List)          // 获取列表
	}
}
