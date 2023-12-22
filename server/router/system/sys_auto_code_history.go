package system

import (
	"github.com/gin-gonic/gin"
	v1 "kirer.cn/server/api/v1"
)

type AutoCodeHistoryRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeHistoryRouter(Router *gin.RouterGroup) {
	router := Router.Group("autoCode")
	api := v1.ApiGroupApp.SystemApiGroup.AutoCodeHistoryApi
	{
		router.POST("delSysHistory", api.Delete) // 删除回滚记录
		router.POST("get", api.Get)              // 根据id获取meta信息
		router.POST("list", api.List)            // 获取回滚记录分页
		router.POST("rollback", api.RollBack)    // 回滚
	}
}
