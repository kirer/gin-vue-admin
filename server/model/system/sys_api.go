package system

import (
	"kirer.cn/server/global"
)

type SysApi struct {
	global.MODEL
	Path   string `json:"path" gorm:"comment:api路径"`             // api路径
	Desc   string `json:"desc" gorm:"comment:api中文描述"`           // api中文描述
	Group  string `json:"group" gorm:"comment:api组"`             // api组
	Method string `json:"method" gorm:"default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

func (SysApi) TableName() string {
	return "sys_api"
}
