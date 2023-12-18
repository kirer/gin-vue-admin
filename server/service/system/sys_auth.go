package system

import (
	"errors"
	"strconv"

	systemReq "kirer.cn/server/model/system/request"

	"gorm.io/gorm"
	"kirer.cn/server/global"
	"kirer.cn/server/model/common/request"
	"kirer.cn/server/model/system"
	"kirer.cn/server/model/system/response"
)

var ErrRoleExistence = errors.New("存在相同角色id")

type AuthService struct{}

var AuthServiceApp = new(AuthService)

func (authService *AuthService) Create(data system.SysAuth) (result system.SysAuth, err error) {
	if err = global.DB.Where("auth_id = ?", data.AuthId).First(&system.SysAuth{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return data, ErrRoleExistence
	}
	e := global.DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&data).Error; err != nil {
			return err
		}
		data.SysBaseMenus = systemReq.DefaultMenu()
		if err = tx.Model(&data).Association("SysBaseMenus").Replace(&data.SysBaseMenus); err != nil {
			return err
		}
		casbinInfos := systemReq.DefaultCasbin()
		authId := strconv.Itoa(int(data.AuthId))
		rules := [][]string{}
		for _, v := range casbinInfos {
			rules = append(rules, []string{authId, v.Path, v.Method})
		}
		return CasbinServiceApp.AddPolicies(tx, rules)
	})
	return data, e
}

func (authService *AuthService) Delete(data *system.SysAuth) error {
	if errors.Is(global.DB.Debug().Preload("Users").First(&data).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(data.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.DB.Where("auth_id = ?", data.AuthId).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.DB.Where("parent_id = ?", data.AuthId).First(&system.SysAuth{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}

	return global.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		if err = tx.Preload("SysBaseMenus").Preload("DataAuthId").Where("auth_id = ?", data.AuthId).First(data).Unscoped().Delete(data).Error; err != nil {
			return err
		}

		if len(data.SysBaseMenus) > 0 {
			if err = tx.Model(data).Association("SysBaseMenus").Delete(data.SysBaseMenus); err != nil {
				return err
			}
		}
		if len(data.DataAuthId) > 0 {
			if err = tx.Model(data).Association("DataAuthId").Delete(data.DataAuthId); err != nil {
				return err
			}
		}
		if err = tx.Delete(&system.SysUserAuth{}, "sys_auth_id = ?", data.AuthId).Error; err != nil {
			return err
		}
		if err = tx.Where("auth_id = ?", data.AuthId).Delete(&[]system.SysAuthBtn{}).Error; err != nil {
			return err
		}
		authId := strconv.Itoa(int(data.AuthId))
		if err = CasbinServiceApp.RemoveFilteredPolicy(tx, authId); err != nil {
			return err
		}
		return nil
	})
}

func (authService *AuthService) Update(data system.SysAuth) (result system.SysAuth, err error) {
	err = global.DB.Where("auth_id = ?", data.AuthId).First(&system.SysAuth{}).Updates(&result).Error
	return
}

func (authService *AuthService) Get(data system.SysAuth) (result system.SysAuth, err error) {
	err = global.DB.Preload("DataAuthId").Where("auth_id = ?", data.AuthId).First(&result).Error
	return
}

func (authService *AuthService) List(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.SysAuth{})
	if err = db.Where("parent_id = ?", "0").Count(&total).Error; total == 0 || err != nil {
		return
	}
	var auth []system.SysAuth
	err = db.Limit(limit).Offset(offset).Preload("DataAuthId").Where("parent_id = ?", "0").Find(&auth).Error
	for k := range auth {
		err = authService.findChildrenAuth(&auth[k])
	}
	return auth, total, err
}

func (authService *AuthService) findChildrenAuth(auth *system.SysAuth) (err error) {
	err = global.DB.Preload("DataAuthId").Where("parent_id = ?", auth.AuthId).Find(&auth.Children).Error
	if len(auth.Children) > 0 {
		for k := range auth.Children {
			err = authService.findChildrenAuth(&auth.Children[k])
		}
	}
	return err
}

func (authService *AuthService) Copy(data response.SysAuthCopyResponse) (auth system.SysAuth, err error) {
	var authBox system.SysAuth
	if !errors.Is(global.DB.Where("auth_id = ?", data.Auth.AuthId).First(&authBox).Error, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}
	data.Auth.Children = []system.SysAuth{}
	menus, err := MenuServiceApp.GetAuth(&request.GetAuthId{AuthId: data.OldAuthId})
	if err != nil {
		return
	}
	var baseMenu []system.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	data.Auth.SysBaseMenus = baseMenu
	err = global.DB.Create(&data.Auth).Error
	if err != nil {
		return
	}
	var btns []system.SysAuthBtn
	err = global.DB.Find(&btns, "auth_id = ?", data.OldAuthId).Error
	if err != nil {
		return
	}
	if len(btns) > 0 {
		for i := range btns {
			btns[i].AuthId = data.Auth.AuthId
		}
		err = global.DB.Create(&btns).Error
		if err != nil {
			return
		}
	}
	paths := CasbinServiceApp.Get(data.OldAuthId)
	err = CasbinServiceApp.Update(data.Auth.AuthId, paths)
	if err != nil {
		_ = authService.Delete(&data.Auth)
	}
	return data.Auth, err
}

func (authService *AuthService) SetData(data system.SysAuth) error {
	var s system.SysAuth
	global.DB.Preload("DataAuthId").First(&s, "auth_id = ?", data.AuthId)
	err := global.DB.Model(&s).Association("DataAuthId").Replace(&data.DataAuthId)
	return err
}

func (authService *AuthService) SetMenu(data *system.SysAuth) error {
	var s system.SysAuth
	global.DB.Preload("SysBaseMenus").First(&s, "auth_id = ?", data.AuthId)
	err := global.DB.Model(&s).Association("SysBaseMenus").Replace(&data.SysBaseMenus)
	return err
}
