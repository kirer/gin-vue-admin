package system

import (
	"errors"
	"fmt"

	"kirer.cn/server/global"
	"kirer.cn/server/model/common/request"
	"kirer.cn/server/model/system"

	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Create
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

type ApiService struct{}

var ApiServiceApp = new(ApiService)

func (apiService *ApiService) Create(data system.SysApi) (err error) {
	if !errors.Is(global.DB.Where("path = ? AND method = ?", data.Path, data.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.DB.Create(&data).Error
}

func (apiService *ApiService) Delete(data system.SysApi) (err error) {
	var entity system.SysApi
	err = global.DB.Where("id = ?", data.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                   // api记录不存在
		return err
	}
	err = global.DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	CasbinServiceApp.ClearCasbin(1, entity.Path, entity.Method)
	if err != nil {
		return err
	}
	return nil
}

func (apiService *ApiService) Deletes(ids request.IdsReq) (err error) {
	var apis []system.SysApi
	err = global.DB.Find(&apis, "id in ?", ids.Ids).Delete(&apis).Error
	if err != nil {
		return
	}
	for _, sysApi := range apis {
		CasbinServiceApp.ClearCasbin(1, sysApi.Path, sysApi.Method)
	}
	if err != nil {
		return
	}
	return err
}

func (apiService *ApiService) Update(data system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.DB.Where("id = ?", data.ID).First(&oldA).Error
	if oldA.Path != data.Path || oldA.Method != data.Method {
		if !errors.Is(global.DB.Where("path = ? AND method = ?", data.Path, data.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return
	}
	err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, data.Path, oldA.Method, data.Method)
	if err != nil {
		return
	}
	err = global.DB.Save(&data).Error
	return
}

func (apiService *ApiService) Get(id int) (result system.SysApi, err error) {
	err = global.DB.Where("id = ?", id).First(&result).Error
	return
}

func (apiService *ApiService) List(data system.SysApi, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.SysApi{})
	if data.Path != "" {
		db = db.Where("path LIKE ?", "%"+data.Path+"%")
	}
	if data.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+data.Desc+"%")
	}
	if data.Method != "" {
		db = db.Where("method = ?", data.Method)
	}
	if data.Group != "" {
		db = db.Where("group = ?", data.Group)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	db = db.Limit(limit).Offset(offset)
	if order != "" {
		var OrderStr string
		orderMap := make(map[string]bool, 5)
		orderMap["id"] = true
		orderMap["path"] = true
		orderMap["group"] = true
		orderMap["desc"] = true
		orderMap["method"] = true
		if orderMap[order] {
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
		} else { // didn't match any order key in `orderMap`
			err = fmt.Errorf("非法的排序字段: %v", order)
			return list, total, err
		}
		err = db.Order(OrderStr).Find(&list).Error
	} else {
		err = db.Order("group").Find(&list).Error
	}
	return list, total, err
}

func (apiService *ApiService) All() (apis []system.SysApi, err error) {
	err = global.DB.Find(&apis).Error
	return
}
