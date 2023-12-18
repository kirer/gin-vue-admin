package core

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"kirer.cn/server/core/internal"
	"kirer.cn/server/global"
	"kirer.cn/server/utils"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("日志创建文件夹:%v\n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))
	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	zap.ReplaceGlobals(logger)
	return logger
}
