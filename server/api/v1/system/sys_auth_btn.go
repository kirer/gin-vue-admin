package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kirer.cn/server/global"
	"kirer.cn/server/model/common/response"
	"kirer.cn/server/model/system/request"
)

type AuthBtnApi struct{}

func (a *AuthBtnApi) Delete(c *gin.Context) {
	id := c.Query("id")
	err := authBtnService.Delete(id)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (a *AuthBtnApi) Update(c *gin.Context) {
	var data request.SysAuthBtnReq
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authBtnService.Update(data)
	if err != nil {
		global.LOG.Error("分配失败!", zap.Error(err))
		response.FailWithMessage("分配失败", c)
		return
	}
	response.OkWithMessage("分配成功", c)
}

func (a *AuthBtnApi) Get(c *gin.Context) {
	var data request.SysAuthBtnReq
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := authBtnService.Get(data)
	if err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(res, "查询成功", c)
}
