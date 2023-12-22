package system

import (
	"kirer.cn/server/global"
)

type JwtBlacklist struct {
	global.MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (JwtBlacklist) TableName() string {
	return "jwt_black"
}