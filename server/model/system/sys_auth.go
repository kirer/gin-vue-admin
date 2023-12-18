package system

import (
	"time"
)

type SysAuth struct {
	CreatedAt     time.Time     // 创建时间
	UpdatedAt     time.Time     // 更新时间
	DeletedAt     *time.Time    `sql:"index"`
	AuthId        uint          `json:"authId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthName      string        `json:"authName" gorm:"comment:角色名"`                                    // 角色名
	ParentId      *uint         `json:"parentId" gorm:"comment:父角色ID"`                                  // 父角色ID
	DataAuthId    []*SysAuth    `json:"dataAuthId" gorm:"many2many:sys_auth_data;"`
	Children      []SysAuth     `json:"children" gorm:"-"`
	SysBaseMenus  []SysBaseMenu `json:"menus" gorm:"many2many:sys_auth_menu;"`
	Users         []SysUser     `json:"-" gorm:"many2many:sys_auth_user;"`
	DefaultRouter string        `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

func (SysAuth) TableName() string {
	return "sys_auth"
}
