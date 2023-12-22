package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kirer.cn/server/global"
	"kirer.cn/server/model/common/response"
	"kirer.cn/server/model/system"
	"kirer.cn/server/model/system/request"
	"kirer.cn/server/utils"
)

type DicDetailApi struct{}

func (s *DicDetailApi) Create(c *gin.Context) {
	var data system.SysDicDetail
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryDetailService.Create(data)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
func (s *DicDetailApi) Delete(c *gin.Context) {
	var data system.SysDicDetail
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryDetailService.Delete(data)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (s *DicDetailApi) Update(c *gin.Context) {
	var data system.SysDicDetail
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryDetailService.Update(&data)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (s *DicDetailApi) Get(c *gin.Context) {
	var data system.SysDicDetail
	err := c.ShouldBindQuery(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(data, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	result, err := dictionaryDetailService.Get(data.ID)
	if err != nil {
		global.LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(result, "查询成功", c)
}

func (s *DicDetailApi) List(c *gin.Context) {
	data, _ := c.GetRawData()
	global.LOG.Error("??", zap.Any(c.Request.RequestURI, string(data)))
	var info request.SysDicDetailSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dictionaryDetailService.List(info)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "获取成功", c)
}
