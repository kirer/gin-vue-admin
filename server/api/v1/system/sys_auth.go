package system

import (
	"kirer.cn/server/global"
	"kirer.cn/server/model/common/request"
	"kirer.cn/server/model/common/response"
	"kirer.cn/server/model/system"
	systemRes "kirer.cn/server/model/system/response"
	"kirer.cn/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthApi struct{}

func (a *AuthApi) Create(c *gin.Context) {
	var data, result system.SysAuth
	var err error
	if err = c.ShouldBindJSON(&data); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err = utils.Verify(data, utils.AuthVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if result, err = authService.Create(data); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
		return
	}
	err = casbinService.FreshCasbin()
	if err != nil {
		global.LOG.Error("创建成功，权限刷新失败。", zap.Error(err))
		response.FailWithMessage("创建成功，权限刷新失败。"+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthResponse{Auth: result}, "创建成功", c)
}

func (a *AuthApi) Delete(c *gin.Context) {
	var data system.SysAuth
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(data, utils.AuthIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 删除角色之前需要判断是否有用户正在使用此角色
	if err := authService.Delete(&data); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
		return
	}
	_ = casbinService.FreshCasbin()
	response.OkWithMessage("删除成功", c)
}

func (a *AuthApi) Update(c *gin.Context) {
	var data system.SysAuth
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(data, utils.AuthVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	result, err := authService.Update(data)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthResponse{Auth: result}, "更新成功", c)
}

func (a *AuthApi) List(c *gin.Context) {
	var info request.PageInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(info, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := authService.List(info)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "获取成功", c)
}

func (a *AuthApi) Copy(c *gin.Context) {
	var data systemRes.SysAuthCopyResponse
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(data, utils.OldAuthVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(data.Auth, utils.AuthVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authBack, err := authService.Copy(data)
	if err != nil {
		global.LOG.Error("拷贝失败!", zap.Error(err))
		response.FailWithMessage("拷贝失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthResponse{Auth: authBack}, "拷贝成功", c)
}

func (a *AuthApi) SetData(c *gin.Context) {
	var auth system.SysAuth
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(auth, utils.AuthIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = authService.SetData(auth)
	if err != nil {
		global.LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
}
