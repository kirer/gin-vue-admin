package system

import (
	"kirer.cn/server/global"
	"kirer.cn/server/model/common/request"
	"kirer.cn/server/model/system"
	systemReq "kirer.cn/server/model/system/request"
)

type RecordService struct{}

func (operationRecordService *RecordService) Create(data system.SysRecord) (err error) {
	err = global.DB.Create(&data).Error
	return err
}

func (operationRecordService *RecordService) Delete(data system.SysRecord) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

func (operationRecordService *RecordService) Deletes(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]system.SysRecord{}, "id in (?)", ids.Ids).Error
	return err
}

func (operationRecordService *RecordService) Get(id uint) (result system.SysRecord, err error) {
	err = global.DB.Where("id = ?", id).First(&result).Error
	return
}

func (operationRecordService *RecordService) List(info systemReq.SysRecordSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&system.SysRecord{})
	var sysRecords []system.SysRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysRecords).Error
	return sysRecords, total, err
}
