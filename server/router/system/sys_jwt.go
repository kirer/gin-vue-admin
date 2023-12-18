package system

import (
	"github.com/gin-gonic/gin"
	v1 "kirer.cn/server/api/v1"
)

type JwtRouter struct{}

func (s *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) {
	router := Router.Group("jwt")
	api := v1.ApiGroupApp.SystemApiGroup.JwtApi
	{
		router.POST("jsonInBlacklist", api.JsonInBlacklist) // jwt加入黑名单
	}
}
