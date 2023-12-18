package system

// SysUserAuth 是 sysUser 和 sysAuth 的连接表
type SysUserAuth struct {
	SysUserId uint `gorm:"column:sys_user_id"`
	SysAuthId uint `gorm:"column:sys_auth_id"`
}

func (s *SysUserAuth) TableName() string {
	return "sys_auth_user"
}
