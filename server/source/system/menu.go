package system

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	. "kirer.cn/server/model/system"
	"kirer.cn/server/source"
)

const initOrderMenu = initOrderAuth + 1

type initMenu struct{}

// auto run
func init() {
	source.RegisterInit(initOrderMenu, &initMenu{})
}

func (i initMenu) InitializerName() string {
	return SysBaseMenu{}.TableName()
}

func (i *initMenu) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, source.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&SysBaseMenu{},
		&SysBaseMenuParameter{},
		&SysBaseMenuBtn{},
	)
}

func (i *initMenu) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&SysBaseMenu{}) &&
		m.HasTable(&SysBaseMenuParameter{}) &&
		m.HasTable(&SysBaseMenuBtn{})
}

func (i *initMenu) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, source.ErrMissingDBContext
	}
	entities := []SysBaseMenu{

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: Meta{Title: "仪表盘", Icon: "odometer"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "admin", Component: "view/admin/index.vue", Sort: 2, Meta: Meta{Title: "管理员", Icon: "user"}},
		{MenuLevel: 0, Hidden: false, ParentId: "2", Path: "auth", Name: "auth", Component: "view/admin/auth/auth.vue", Sort: 1, Meta: Meta{Title: "角色", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: "2", Path: "user", Name: "user", Component: "view/admin/user/user.vue", Sort: 2, Meta: Meta{Title: "用户", Icon: "coordinate"}},
		{MenuLevel: 0, Hidden: false, ParentId: "2", Path: "menu", Name: "menu", Component: "view/admin/menu/menu.vue", Sort: 3, Meta: Meta{Title: "菜单", Icon: "tickets", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "2", Path: "api", Name: "api", Component: "view/admin/api/api.vue", Sort: 4, Meta: Meta{Title: "api", Icon: "platform", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "2", Path: "dic", Name: "dic", Component: "view/admin/dic/dic.vue", Sort: 5, Meta: Meta{Title: "字典", Icon: "notebook"}},
		{MenuLevel: 0, Hidden: true, ParentId: "2", Path: "dic_detail", Name: "dic_detail", Component: "view/admin/dic/dic_detail.vue", Sort: 5, Meta: Meta{Title: "字典详情", Icon: "notebook"}},
		{MenuLevel: 0, Hidden: false, ParentId: "2", Path: "record", Name: "record", Component: "view/admin/record/record.vue", Sort: 6, Meta: Meta{Title: "记录", Icon: "pie-chart"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "tools", Name: "tools", Component: "view/tools/index.vue", Sort: 3, Meta: Meta{Title: "工具", Icon: "tools"}},
		{MenuLevel: 0, Hidden: false, ParentId: "10", Path: "auto_package", Name: "auto_package", Component: "view/tools/auto/package.vue", Sort: 1, Meta: Meta{Title: "Package", Icon: "folder"}},
		{MenuLevel: 0, Hidden: false, ParentId: "10", Path: "auto_creator", Name: "auto_creator", Component: "view/tools/auto/creator.vue", Sort: 2, Meta: Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "10", Path: "auto_manager", Name: "auto_manager", Component: "view/tools/auto/manager.vue", Sort: 3, Meta: Meta{Title: "自动化代码管理", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: true, ParentId: "10", Path: "auto_creator/:id", Name: "auto_creator", Component: "view/tools/auto/creator.vue", Sort: 0, Meta: Meta{Title: "自动化代码-${id}", Icon: "magic-stick"}},
		{MenuLevel: 0, Hidden: false, ParentId: "10", Path: "config", Name: "config", Component: "view/tools/config.vue", Sort: 4, Meta: Meta{Title: "系统配置", Icon: "operation"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "plugin", Name: "plugin", Component: "view/routerHolder.vue", Sort: 4, Meta: Meta{Title: "插件系统", Icon: "cherry"}},
		{MenuLevel: 0, Hidden: false, ParentId: "17", Path: "plugin_install", Name: "plugin_install", Component: "view/plugin/install.vue", Sort: 1, Meta: Meta{Title: "插件安装", Icon: "box"}},
		{MenuLevel: 0, Hidden: false, ParentId: "17", Path: "plugin_template", Name: "plugin_template", Component: "view/system/plugin/template.vue", Sort: 2, Meta: Meta{Title: "插件模板", Icon: "folder"}},
		{MenuLevel: 0, Hidden: false, ParentId: "17", Path: "plugin_package", Name: "plugin_package", Component: "view/system/plugin/package.vue", Sort: 3, Meta: Meta{Title: "打包插件", Icon: "files"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 5, Meta: Meta{Title: "服务器状态", Icon: "cloudy"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 6, Meta: Meta{Title: "关于", Icon: "info-filled"}},
		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 7, Meta: Meta{Title: "个人信息", Icon: "message"}},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, SysBaseMenu{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initMenu) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ?", "admin").First(&SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
