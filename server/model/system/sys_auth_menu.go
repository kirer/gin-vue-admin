package system

type SysMenu struct {
	SysBaseMenu
	MenuId     string                 `json:"menuId" gorm:"comment:菜单ID"`
	AuthId     uint                   `json:"-" gorm:"comment:角色ID"`
	Children   []SysMenu              `json:"children" gorm:"-"`
	Parameters []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
	Btns       map[string]uint        `json:"btns" gorm:"-"`
}

type SysAuthMenu struct {
	MenuId string `json:"menuId" gorm:"comment:菜单ID;column:sys_base_menu_id"`
	AuthId string `json:"-" gorm:"comment:角色ID;column:sys_auth_id"`
}

func (s SysAuthMenu) TableName() string {
	return "sys_auth_menu"
}
