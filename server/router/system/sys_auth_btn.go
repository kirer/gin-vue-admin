package system

import (
	"github.com/gin-gonic/gin"
	v1 "kirer.cn/server/api/v1"
)

type AuthBtnRouter struct{}

func (s *AuthBtnRouter) InitAuthBtnRouterRouter(Router *gin.RouterGroup) {
	//router := Router.Group("authBtn").Use(middleware.Record())
	router2 := Router.Group("auth_btn")
	api := v1.ApiGroupApp.SystemApiGroup.AuthBtnApi
	{
		router2.DELETE("delete", api.Delete)
		router2.PUT("update", api.Update)
		router2.POST("get", api.Get)
	}
}
