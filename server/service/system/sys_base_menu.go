package system

import (
	"errors"

	"gorm.io/gorm"
	"kirer.cn/server/global"
	"kirer.cn/server/model/system"
)

type BaseMenuService struct{}

func (baseMenuService *BaseMenuService) Delete(id int) (err error) {
	err = global.DB.Preload("MenuBtn").Preload("Parameters").Where("parent_id = ?", id).First(&system.SysBaseMenu{}).Error
	if err != nil {
		var menu system.SysBaseMenu
		db := global.DB.Preload("SysAuths").Where("id = ?", id).First(&menu).Delete(&menu)
		err = global.DB.Delete(&system.SysBaseMenuParameter{}, "sys_menu_id = ?", id).Error
		if err != nil {
			return
		}
		err = global.DB.Delete(&system.SysBaseMenuBtn{}, "sys_menu_id = ?", id).Error
		if err != nil {
			return
		}
		err = global.DB.Delete(&system.SysAuthBtn{}, "sys_menu_id = ?", id).Error
		if err != nil {
			return
		}
		if len(menu.SysAuths) > 0 {
			err = global.DB.Model(&menu).Association("SysAuths").Delete(&menu.SysAuths)
		} else {
			err = db.Error
		}
		if err != nil {
			return
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return
}

func (baseMenuService *BaseMenuService) Update(menu system.SysBaseMenu) (err error) {
	var oldMenu system.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["close_tab"] = menu.CloseTab
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["active_name"] = menu.ActiveName
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				global.LOG.Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}
		txErr := tx.Unscoped().Delete(&system.SysBaseMenuParameter{}, "sys_menu_id = ?", menu.ID).Error
		if txErr != nil {
			global.LOG.Debug(txErr.Error())
			return txErr
		}
		txErr = tx.Unscoped().Delete(&system.SysBaseMenuBtn{}, "sys_menu_id = ?", menu.ID).Error
		if txErr != nil {
			global.LOG.Debug(txErr.Error())
			return txErr
		}
		if len(menu.Parameters) > 0 {
			for k := range menu.Parameters {
				menu.Parameters[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.Parameters).Error
			if txErr != nil {
				global.LOG.Debug(txErr.Error())
				return txErr
			}
		}

		if len(menu.MenuBtn) > 0 {
			for k := range menu.MenuBtn {
				menu.MenuBtn[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.MenuBtn).Error
			if txErr != nil {
				global.LOG.Debug(txErr.Error())
				return txErr
			}
		}

		txErr = db.Updates(upDateMap).Error
		if txErr != nil {
			global.LOG.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

func (baseMenuService *BaseMenuService) Get(id int) (menu system.SysBaseMenu, err error) {
	err = global.DB.Preload("MenuBtn").Preload("Parameters").Where("id = ?", id).First(&menu).Error
	return
}
