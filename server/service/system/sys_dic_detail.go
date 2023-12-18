package system

import (
	"kirer.cn/server/global"
	"kirer.cn/server/model/system"
	"kirer.cn/server/model/system/request"
)

type DicDetailService struct{}

func (DicDetailService *DicDetailService) Create(data system.SysDicDetail) (err error) {
	err = global.DB.Create(&data).Error
	return err
}

func (DicDetailService *DicDetailService) Delete(data system.SysDicDetail) (err error) {
	err = global.DB.Delete(&data).Error
	return err
}

func (DicDetailService *DicDetailService) Update(data *system.SysDicDetail) (err error) {
	err = global.DB.Save(data).Error
	return err
}

func (DicDetailService *DicDetailService) Get(id uint) (result system.SysDicDetail, err error) {
	err = global.DB.Where("id = ?", id).First(&result).Error
	return
}

func (DicDetailService *DicDetailService) List(info request.SysDicDetailSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&system.SysDicDetail{})
	var dicDetails []system.SysDicDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.DicID != 0 {
		db = db.Where("dic_id = ?", info.DicID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort").Find(&dicDetails).Error
	return dicDetails, total, err
}
