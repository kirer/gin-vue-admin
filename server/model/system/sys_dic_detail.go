package system

import (
	"kirer.cn/server/global"
)

type SysDicDetail struct {
	global.MODEL
	Label  string `json:"label" form:"label" gorm:"column:label;comment:展示值"`       // 展示值
	Value  int    `json:"value" form:"value" gorm:"column:value;comment:字典值"`       // 字典值
	Extend string `json:"extend" form:"extend" gorm:"column:extend;comment:扩展值"`    // 扩展值
	Status *bool  `json:"status" form:"status" gorm:"column:status;comment:启用状态"`   // 启用状态
	Sort   int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序标记"`         // 排序标记
	DicID  int    `json:"dicID" form:"dicID" gorm:"column:sys_dic_id;comment:关联标记"` // 关联标记
}

func (SysDicDetail) TableName() string {
	return "sys_dic_detail"
}
