package system

import (
	"errors"

	"gorm.io/gorm"
	"kirer.cn/server/global"
	"kirer.cn/server/model/system"
	"kirer.cn/server/model/system/request"
	"kirer.cn/server/model/system/response"
)

type AuthBtnService struct{}

func (a *AuthBtnService) Delete(data string) (err error) {
	fErr := global.DB.First(&system.SysAuthBtn{}, "sys_base_menu_btn_id = ?", data).Error
	if errors.Is(fErr, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.New("此按钮正在被使用无法删除")
}

func (a *AuthBtnService) Update(data request.SysAuthBtnReq) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var authBtn []system.SysAuthBtn
		err = tx.Delete(&[]system.SysAuthBtn{}, "auth_id = ? and sys_menu_id = ?", data.AuthId, data.MenuID).Error
		if err != nil {
			return err
		}
		for _, v := range data.Selected {
			authBtn = append(authBtn, system.SysAuthBtn{
				AuthId:           data.AuthId,
				SysMenuID:        data.MenuID,
				SysBaseMenuBtnID: v,
			})
		}
		if len(authBtn) > 0 {
			err = tx.Create(&authBtn).Error
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (a *AuthBtnService) Get(data request.SysAuthBtnReq) (result response.SysAuthBtnRes, err error) {
	var authBtn []system.SysAuthBtn
	err = global.DB.Find(&authBtn, "auth_id = ? and sys_menu_id = ?", data.AuthId, data.MenuID).Error
	if err != nil {
		return
	}
	var selected []uint
	for _, v := range authBtn {
		selected = append(selected, v.SysBaseMenuBtnID)
	}
	result.Selected = selected
	return result, err
}
