package system

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	sysModel "kirer.cn/server/model/system"
	"kirer.cn/server/source"
)

const initOrderMenuAuth = initOrderMenu + initOrderAuth

type initMenuAuth struct{}

// auto run
func init() {
	source.RegisterInit(initOrderMenuAuth, &initMenuAuth{})
}

func (i *initMenuAuth) MigrateTable(ctx context.Context) (context.Context, error) {
	return ctx, nil // do nothing
}

func (i *initMenuAuth) TableCreated(ctx context.Context) bool {
	return false // always replace
}

func (i initMenuAuth) InitializerName() string {
	return sysModel.SysAuthMenu{}.TableName()
}

func (i *initMenuAuth) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, source.ErrMissingDBContext
	}
	authorities, ok := ctx.Value(initAuth{}.InitializerName()).([]sysModel.SysAuth)
	if !ok {
		return ctx, errors.Wrap(source.ErrMissingDependentContext, "创建 [菜单-权限] 关联失败, 未找到权限表初始化数据")
	}
	menus, ok := ctx.Value(initMenu{}.InitializerName()).([]sysModel.SysBaseMenu)
	if !ok {
		return next, errors.Wrap(errors.New(""), "创建 [菜单-权限] 关联失败, 未找到菜单表初始化数据")
	}
	next = ctx
	// 888
	if err = db.Model(&authorities[0]).Association("SysBaseMenus").Replace(menus); err != nil {
		return next, err
	}
	return next, nil
}

func (i *initMenuAuth) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	auth := &sysModel.SysAuth{}
	if ret := db.Model(auth).Where("auth_id = ?", 888).Preload("SysBaseMenus").Find(auth); ret != nil {
		if ret.Error != nil {
			return false
		}
		return len(auth.SysBaseMenus) > 0
	}
	return false
}
