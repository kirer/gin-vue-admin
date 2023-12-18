package core

import (
	"fmt"
	"path/filepath"

	"kirer.cn/server/core/internal"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"kirer.cn/server/global"
	_ "kirer.cn/server/packfile"
)

func Viper(path ...string) *viper.Viper {
	var config = internal.ConfigDefaultFile
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取错误:%s", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件变更:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		panic(fmt.Errorf("配置文件解析错误:%s", err))
	}
	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	global.CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
