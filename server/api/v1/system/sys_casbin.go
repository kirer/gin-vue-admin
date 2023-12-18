package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kirer.cn/server/global"
	"kirer.cn/server/model/common/response"
	"kirer.cn/server/model/system/request"
	systemRes "kirer.cn/server/model/system/response"
	"kirer.cn/server/utils"
)

type CasbinApi struct{}

func (cas *CasbinApi) Update(c *gin.Context) {
	var data request.CasbinInReceive
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(data, utils.AuthIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = casbinService.Update(data.AuthId, data.CasbinInfos)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (cas *CasbinApi) Get(c *gin.Context) {
	var data request.CasbinInReceive
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(data, utils.AuthIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	paths := casbinService.Get(data.AuthId)
	response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "获取成功", c)
}
