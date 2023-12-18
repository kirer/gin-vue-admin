package system

import (
	"errors"

	"gorm.io/gorm"
	"kirer.cn/server/global"
	"kirer.cn/server/model/system"
)

type DicService struct{}

func (dictionaryService *DicService) Create(data system.SysDic) (err error) {
	if (!errors.Is(global.DB.First(&system.SysDic{}, "type = ?", data.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = global.DB.Create(&data).Error
	return err
}

func (dictionaryService *DicService) Delete(data system.SysDic) (err error) {
	err = global.DB.Where("id = ?", data.ID).Preload("SysDicDetails").First(&data).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("请不要搞事")
	}
	if err != nil {
		return err
	}
	err = global.DB.Delete(&data).Error
	if err != nil {
		return err
	}
	if data.Details != nil {
		return global.DB.Where("sys_dic_id=?", data.ID).Delete(data.Details).Error
	}
	return
}

func (dictionaryService *DicService) Update(data *system.SysDic) (err error) {
	var dict system.SysDic
	sysDicMap := map[string]interface{}{
		"Name":   data.Name,
		"Type":   data.Type,
		"Status": data.Status,
		"Desc":   data.Desc,
	}
	db := global.DB.Where("id = ?", data.ID).First(&dict)
	if dict.Type != data.Type {
		if !errors.Is(global.DB.First(&system.SysDic{}, "type = ?", data.Type).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的type，不允许创建")
		}
	}
	err = db.Updates(sysDicMap).Error
	return err
}

func (dictionaryService *DicService) Get(Type string, Id uint, status *bool) (sysDic system.SysDic, err error) {
	var flag = false
	if status == nil {
		flag = true
	} else {
		flag = *status
	}
	err = global.DB.Where("(type = ? OR id = ?) and status = ?", Type, Id, flag).Preload("Details", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", true).Order("sort")
	}).First(&sysDic).Error
	return
}

func (dictionaryService *DicService) List() (list interface{}, err error) {
	var sysDics []system.SysDic
	err = global.DB.Find(&sysDics).Error
	return sysDics, err
}
