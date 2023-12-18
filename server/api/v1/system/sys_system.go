package system

import (
	"kirer.cn/server/global"
	"kirer.cn/server/model/common/response"
	"kirer.cn/server/model/system"
	systemRes "kirer.cn/server/model/system/response"
	"kirer.cn/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApi struct{}

func (s *SystemApi) GetConfig(c *gin.Context) {
	config, err := systemConfigService.GetConfig()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "获取成功", c)
}

func (s *SystemApi) SetConfig(c *gin.Context) {
	var data system.System
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = systemConfigService.SetConfig(data)
	if err != nil {
		global.LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

func (s *SystemApi) GetInfo(c *gin.Context) {
	server, err := systemConfigService.GetInfo()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
}

func (s *SystemApi) Reload(c *gin.Context) {
	err := utils.Reload()
	if err != nil {
		global.LOG.Error("重启系统失败!", zap.Error(err))
		response.FailWithMessage("重启系统失败", c)
		return
	}
	response.OkWithMessage("重启系统成功", c)
}
