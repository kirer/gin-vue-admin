package system

import (
	"github.com/gin-gonic/gin"
	v1 "kirer.cn/server/api/v1"
	"kirer.cn/server/middleware"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	router := Router.Group("casbin").Use(middleware.Record())
	router2 := Router.Group("casbin")
	api := v1.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		router.PUT("update", api.Update)
	}
	{
		router2.POST("get", api.Get)
	}
}
