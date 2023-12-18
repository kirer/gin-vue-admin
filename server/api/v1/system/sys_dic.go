package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kirer.cn/server/global"
	"kirer.cn/server/model/common/response"
	"kirer.cn/server/model/system"
)

type DicApi struct{}

func (s *DicApi) Create(c *gin.Context) {
	var data system.SysDic
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.Create(data)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (s *DicApi) Delete(c *gin.Context) {
	var data system.SysDic
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.Delete(data)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (s *DicApi) Update(c *gin.Context) {
	var data system.SysDic
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.Update(&data)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (s *DicApi) Get(c *gin.Context) {
	var data system.SysDic
	err := c.ShouldBindQuery(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysDic, err := dictionaryService.Get(data.Type, data.ID, data.Status)
	if err != nil {
		global.LOG.Error("字典未创建或未开启!", zap.Error(err))
		response.FailWithMessage("字典未创建或未开启", c)
		return
	}
	response.OkWithDetailed(gin.H{"resysDic": sysDic}, "查询成功", c)
}

func (s *DicApi) List(c *gin.Context) {
	list, err := dictionaryService.List()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(list, "获取成功", c)
}
