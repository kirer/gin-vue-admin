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
	SysBaseMenus  []SysBaseMenu `json:"menus" gorm:"many2many:sys_auth_menu;"`
	Users         []SysUser     `json:"-" gorm:"many2many:sys_auth_user;"`
	Children      []SysAuth     `json:"children" gorm:"-"`
	DefaultRouter string        `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

func (SysAuth) TableName() string {
	return "sys_auth"
}

type SysAuthUser struct {
	SysUserId uint `gorm:"column:sys_user_id;comment:用户ID;"`
	SysAuthId uint `gorm:"column:sys_auth_auth_id;comment:角色ID;"`
}

func (SysAuthUser) TableName() string {
	return "sys_auth_user"
}

type SysAuthMenu struct {
	MenuId string `gorm:"column:sys_base_menu_id;comment:菜单ID;"`
	AuthId string `gorm:"column:sys_auth_auth_id;comment:角色ID;"`
}

func (SysAuthMenu) TableName() string {
	return "sys_auth_menu"
}

type SysAuthBtn struct {
	AuthId           uint           `gorm:"column:auth_id;comment:角色ID"`
	SysMenuID        uint           `gorm:"column:sys_menu_id;comment:菜单ID"`
	SysBaseMenuBtnID uint           `gorm:"column:sys_base_menu_btn_id;comment:按钮ID"`
	SysBaseMenuBtn   SysBaseMenuBtn `gorm:"comment:按钮详情"`
}

func (SysAuthBtn) TableName() string {
	return "sys_auth_btn"
}
