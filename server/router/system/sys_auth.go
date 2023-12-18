package system

import (
	"github.com/gin-gonic/gin"
	v1 "kirer.cn/server/api/v1"
	"kirer.cn/server/middleware"
)

type AuthRouter struct{}

func (s *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	router := Router.Group("auth").Use(middleware.Record())
	router2 := Router.Group("auth")
	api := v1.ApiGroupApp.SystemApiGroup.AuthApi
	{
		router.POST("create", api.Create)   // 创建角色
		router.DELETE("delete", api.Delete) // 删除角色
		router.PUT("update", api.Update)    // 更新角色
		router.POST("copy", api.Copy)       // 拷贝角色
		router.POST("set_data", api.SetData) // 设置角色资源权限
	}
	{
		router2.POST("list", api.List) // 获取角色列表
	}
}
