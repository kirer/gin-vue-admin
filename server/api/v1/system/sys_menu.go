package system

import (
	"kirer.cn/server/global"
	"kirer.cn/server/model/common/request"
	"kirer.cn/server/model/common/response"
	"kirer.cn/server/model/system"
	systemReq "kirer.cn/server/model/system/request"
	systemRes "kirer.cn/server/model/system/response"
	"kirer.cn/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthMenuApi struct{}

func (a *AuthMenuApi) Create(c *gin.Context) {
	var data system.SysBaseMenu
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(data, utils.MenuVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(data.Meta, utils.MenuMetaVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = menuService.Create(data)
	if err != nil {
		global.LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

func (a *AuthMenuApi) CreateAuth(c *gin.Context) {
	var data systemReq.AddMenuAuthInfo
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(data, utils.AuthIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.CreateAuth(data.Menus, data.AuthId); err != nil {
		global.LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

func (a *AuthMenuApi) Delete(c *gin.Context) {
	var data request.GetById
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(data, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = baseMenuService.Delete(data.ID)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (a *AuthMenuApi) Update(c *gin.Context) {
	var data system.SysBaseMenu
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(data, utils.MenuVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(data.Meta, utils.MenuMetaVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = baseMenuService.Update(data)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (a *AuthMenuApi) Get(c *gin.Context) {
	var data request.GetById
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(data, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menu, err := baseMenuService.Get(data.ID)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
}

func (a *AuthMenuApi) GetAuth(c *gin.Context) {
	var data request.GetAuthId
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
	global.LOG.Error(">>>", zap.Any(">>>", data))
	menus, err := menuService.GetAuth(&data)
	global.LOG.Error(">>>", zap.Any(">>>>", menus))
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"menus": menus}, "获取成功", c)
}

func (a *AuthMenuApi) GetCurrent(c *gin.Context) {
	menus, err := menuService.GetByAuth(utils.GetUserAuthId(c))
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if menus == nil {
		menus = []system.SysMenu{}
	}
	response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
}

func (a *AuthMenuApi) List(c *gin.Context) {
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
	menuList, total, err := menuService.List()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     menuList,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "获取成功", c)
}

func (a *AuthMenuApi) All(c *gin.Context) {
	menus, err := menuService.All()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysBaseMenusResponse{Menus: menus}, "获取成功", c)
}
