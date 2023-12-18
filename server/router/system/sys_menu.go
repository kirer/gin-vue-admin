package system

import (
	"github.com/gin-gonic/gin"
	v1 "kirer.cn/server/api/v1"
	"kirer.cn/server/middleware"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	router := Router.Group("menu").Use(middleware.Record())
	router2 := Router.Group("menu")
	api := v1.ApiGroupApp.SystemApiGroup.AuthMenuApi
	{
		router.POST("create", api.Create)          // 新增菜单
		router.POST("create_auth", api.CreateAuth) //	增加menu和角色关联关系
		router.DELETE("delete", api.Delete)        // 删除菜单
		router.PUT("update", api.Update)           // 更新菜单
	}
	{
		router2.POST("get", api.Get)                // 根据id获取菜单
		router2.POST("get_auth", api.GetAuth)       // 获取指定角色menu
		router2.POST("get_current", api.GetCurrent) // 获取菜单树
		router2.POST("list", api.List)              // 分页获取基础menu列表
		router2.POST("all", api.All)                // 获取用户动态路由
	}
}
