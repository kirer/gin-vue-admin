package system

import (
	"errors"
	"strconv"

	"gorm.io/gorm"
	"kirer.cn/server/global"
	"kirer.cn/server/model/common/request"
	"kirer.cn/server/model/system"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

func (menuService *MenuService) Create(data system.SysBaseMenu) error {
	if !errors.Is(global.DB.Where("name = ?", data.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.DB.Create(&data).Error
}

func (menuService *MenuService) CreateAuth(menus []system.SysBaseMenu, authId uint) (err error) {
	var auth system.SysAuth
	auth.AuthId = authId
	auth.SysBaseMenus = menus
	err = AuthServiceApp.SetMenu(&auth)
	return err
}

func (menuService *MenuService) GetAuth(info *request.GetAuthId) (menus []system.SysMenu, err error) {
	var baseMenu []system.SysBaseMenu
	var SysAuthMenus []system.SysAuthMenu
	err = global.DB.Where("sys_auth_id = ?", info.AuthId).Find(&SysAuthMenus).Error
	if err != nil {
		return
	}

	var MenuIds []string

	for i := range SysAuthMenus {
		MenuIds = append(MenuIds, SysAuthMenus[i].MenuId)
	}

	err = global.DB.Where("id in (?) ", MenuIds).Order("sort").Find(&baseMenu).Error

	for i := range baseMenu {
		menus = append(menus, system.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthId:      info.AuthId,
			MenuId:      strconv.Itoa(int(baseMenu[i].ID)),
			Parameters:  baseMenu[i].Parameters,
		})
	}
	return menus, err
}

func (menuService *MenuService) GetByAuth(authId uint) (menus []system.SysMenu, err error) {
	menuTree, err := menuService.getMenuTreeMap(authId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

func (menuService *MenuService) List() (list interface{}, total int64, err error) {
	var menuList []system.SysBaseMenu
	treeMap, err := menuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return menuList, total, err
}

func (menuService *MenuService) All() (menus []system.SysBaseMenu, err error) {
	treeMap, err := menuService.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return menus, err
}

func (menuService *MenuService) UserAuthDefaultRouter(user *system.SysUser) {
	var menuIds []string
	err := global.DB.Model(&system.SysAuthMenu{}).Where("sys_auth_id = ?", user.AuthId).Pluck("sys_base_menu_id", &menuIds).Error
	if err != nil {
		return
	}
	var am system.SysBaseMenu
	err = global.DB.First(&am, "name = ? and id in (?)", user.Auth.DefaultRouter, menuIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Auth.DefaultRouter = "404"
	}
}

func (menuService *MenuService) getMenuTreeMap(authId uint) (treeMap map[string][]system.SysMenu, err error) {
	var allMenus []system.SysMenu
	var baseMenu []system.SysBaseMenu
	var btns []system.SysAuthBtn
	treeMap = make(map[string][]system.SysMenu)

	var SysAuthMenus []system.SysAuthMenu
	err = global.DB.Where("sys_auth_id = ?", authId).Find(&SysAuthMenus).Error
	if err != nil {
		return
	}

	var MenuIds []string

	for i := range SysAuthMenus {
		MenuIds = append(MenuIds, SysAuthMenus[i].MenuId)
	}

	err = global.DB.Where("id in (?)", MenuIds).Order("sort").Preload("Parameters").Find(&baseMenu).Error
	if err != nil {
		return
	}

	for i := range baseMenu {
		allMenus = append(allMenus, system.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthId:      authId,
			MenuId:      strconv.Itoa(int(baseMenu[i].ID)),
			Parameters:  baseMenu[i].Parameters,
		})
	}

	err = global.DB.Where("auth_id = ?", authId).Preload("SysBaseMenuBtn").Find(&btns).Error
	if err != nil {
		return
	}
	var btnMap = make(map[uint]map[string]uint)
	for _, v := range btns {
		if btnMap[v.SysMenuID] == nil {
			btnMap[v.SysMenuID] = make(map[string]uint)
		}
		btnMap[v.SysMenuID][v.SysBaseMenuBtn.Name] = authId
	}
	for _, v := range allMenus {
		v.Btns = btnMap[v.ID]
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

func (menuService *MenuService) getChildrenList(menu *system.SysMenu, treeMap map[string][]system.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (menuService *MenuService) getBaseChildrenList(menu *system.SysBaseMenu, treeMap map[string][]system.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (menuService *MenuService) getBaseMenuTreeMap() (treeMap map[string][]system.SysBaseMenu, err error) {
	var allMenus []system.SysBaseMenu
	treeMap = make(map[string][]system.SysBaseMenu)
	err = global.DB.Order("sort").Preload("MenuBtn").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}
