package main

import (
	"fmt"

	"kirer.cn/server/core"
	"kirer.cn/server/global"
	"kirer.cn/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	global.VP = core.Viper()      // 初始化Viper
	global.LOG = core.Zap()       // 初始化zap日志库
	global.DB = initialize.Gorm() // gorm连接数据库
	db, err := global.DB.DB()
	if err != nil {
		panic(fmt.Errorf("数据库连接失败:%s", err))
	}
	defer db.Close()
	initialize.Timer()
	initialize.OtherInit()
	core.RunWindowsServer()
}
