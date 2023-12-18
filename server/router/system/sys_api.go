package system

import (
	"github.com/gin-gonic/gin"
	v1 "kirer.cn/server/api/v1"
	"kirer.cn/server/middleware"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	router := Router.Group("api").Use(middleware.Record())
	router2 := Router.Group("api")
	router3 := RouterPub.Group("api")
	api := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		router.POST("create", api.Create)     // 创建Api
		router.DELETE("delete", api.Delete)   // 删除Api
		router.DELETE("deletes", api.Deletes) // 删除选中api
		router.PUT("update", api.Update)     // 更新api
		router.POST("get", api.Get)           // 获取单条Api消息
	}
	{
		router2.POST("list", api.List) // 获取Api列表
		router2.POST("all", api.All)   // 获取所有api
	}
	{
		router3.GET("freshCasbin", api.FreshCasbin) // 刷新casbin权限
	}
}
