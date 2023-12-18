package system

import (
	"go.uber.org/zap"
	"kirer.cn/server/config"
	"kirer.cn/server/global"
	"kirer.cn/server/model/system"
	"kirer.cn/server/utils"
)

type SystemConfigService struct{}

func (systemConfigService *SystemConfigService) GetConfig() (conf config.Server, err error) {
	return global.CONFIG, nil
}

func (systemConfigService *SystemConfigService) SetConfig(data system.System) (err error) {
	cs := utils.StructToMap(data.Config)
	for k, v := range cs {
		global.VP.Set(k, v)
	}
	err = global.VP.WriteConfig()
	return err
}

func (systemConfigService *SystemConfigService) GetInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); err != nil {
		global.LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
