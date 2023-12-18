package system

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	sysModel "kirer.cn/server/model/system"
	"kirer.cn/server/source"
)

type initApi struct{}

const initOrderApi = source.InitOrderSystem + 1

// auto run
func init() {
	source.RegisterInit(initOrderApi, &initApi{})
}

func (i initApi) InitializerName() string {
	return sysModel.SysApi{}.TableName()
}

func (i *initApi) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, source.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysApi{})
}

func (i *initApi) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysApi{})
}

func (i *initApi) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, source.ErrMissingDBContext
	}
	entities := []sysModel.SysApi{
		{Group: "jwt", Method: "POST", Path: "/jwt/jsonInBlacklist", Desc: "jwt加入黑名单(退出，必选)"},

		{Group: "系统用户", Method: "DELETE", Path: "/user/deleteUser", Desc: "删除用户"},
		{Group: "系统用户", Method: "POST", Path: "/user/admin_register", Desc: "用户注册"},
		{Group: "系统用户", Method: "POST", Path: "/user/getUserList", Desc: "获取用户列表"},
		{Group: "系统用户", Method: "PUT", Path: "/user/setUserInfo", Desc: "设置用户信息"},
		{Group: "系统用户", Method: "PUT", Path: "/user/setSelfInfo", Desc: "设置自身信息(必选)"},
		{Group: "系统用户", Method: "GET", Path: "/user/getUserInfo", Desc: "获取自身信息(必选)"},
		{Group: "系统用户", Method: "POST", Path: "/user/setUserAuthorities", Desc: "设置权限组"},
		{Group: "系统用户", Method: "POST", Path: "/user/changePassword", Desc: "修改密码（建议选择)"},
		{Group: "系统用户", Method: "POST", Path: "/user/setUserAuth", Desc: "修改用户角色(必选)"},
		{Group: "系统用户", Method: "POST", Path: "/user/resetPassword", Desc: "重置用户密码"},

		{Group: "api", Method: "POST", Path: "/api/create", Desc: "创建api"},
		{Group: "api", Method: "DELETE", Path: "/api/delete", Desc: "删除Api"},
		{Group: "api", Method: "DELETE", Path: "/api/deletes", Desc: "批量删除api"},
		{Group: "api", Method: "PUT", Path: "/api/update", Desc: "更新Api"},
		{Group: "api", Method: "POST", Path: "/api/get", Desc: "获取api详细信息"},
		{Group: "api", Method: "POST", Path: "/api/list", Desc: "获取api列表"},
		{Group: "api", Method: "POST", Path: "/api/all", Desc: "获取所有api"},

		{Group: "角色", Method: "POST", Path: "/auth/create", Desc: "创建角色"},
		{Group: "角色", Method: "DELETE", Path: "/auth/delete", Desc: "删除角色"},
		{Group: "角色", Method: "PUT", Path: "/auth/update", Desc: "更新角色信息"},
		{Group: "角色", Method: "POST", Path: "/auth/list", Desc: "获取角色列表"},
		{Group: "角色", Method: "POST", Path: "/auth/copy", Desc: "拷贝角色"},
		{Group: "角色", Method: "POST", Path: "/auth/set_data", Desc: "设置角色资源权限"},

		{Group: "按钮权限", Method: "DELETE", Path: "/auth_btn/delete", Desc: "删除按钮"},
		{Group: "按钮权限", Method: "PUT", Path: "/auth_btn/update", Desc: "设置按钮权限"},
		{Group: "按钮权限", Method: "POST", Path: "/auth_btn/get", Desc: "获取已有按钮权限"},

		{Group: "Casbin", Method: "PUT", Path: "/casbin/update", Desc: "更改角色api权限"},
		{Group: "Casbin", Method: "POST", Path: "/casbin/get", Desc: "获取权限列表"},

		{Group: "菜单", Method: "POST", Path: "/menu/create", Desc: "新增菜单"},
		{Group: "菜单", Method: "POST", Path: "/menu/create_auth", Desc: "增加menu和角色关联关系"},
		{Group: "菜单", Method: "DELETE", Path: "/menu/delete", Desc: "删除菜单"},
		{Group: "菜单", Method: "PUT", Path: "/menu/update", Desc: "更新菜单"},
		{Group: "菜单", Method: "POST", Path: "/menu/get", Desc: "根据id获取菜单"},
		{Group: "菜单", Method: "POST", Path: "/menu/get_auth", Desc: "获取指定角色menu"},
		{Group: "菜单", Method: "POST", Path: "/menu/get_current", Desc: "获取菜单树(必选)"},
		{Group: "菜单", Method: "POST", Path: "/menu/list", Desc: "分页获取基础menu列表"},
		{Group: "菜单", Method: "POST", Path: "/menu/all", Desc: "获取用户动态路由"},

		{Group: "系统字典", Method: "POST", Path: "/dic/create", Desc: "新增字典"},
		{Group: "系统字典", Method: "DELETE", Path: "/dic/delete", Desc: "删除字典"},
		{Group: "系统字典", Method: "PUT", Path: "/dic/update", Desc: "更新字典"},
		{Group: "系统字典", Method: "GET", Path: "/dic/get", Desc: "根据ID获取字典"},
		{Group: "系统字典", Method: "GET", Path: "/dic/list", Desc: "获取字典列表"},

		{Group: "系统字典详情", Method: "POST", Path: "/dic_detail/create", Desc: "新增字典内容"},
		{Group: "系统字典详情", Method: "DELETE", Path: "/dic_detail/delete", Desc: "删除字典内容"},
		{Group: "系统字典详情", Method: "PUT", Path: "/dic_detail/update", Desc: "更新字典内容"},
		{Group: "系统字典详情", Method: "GET", Path: "/dic_detail/get", Desc: "根据ID获取字典内容"},
		{Group: "系统字典详情", Method: "GET", Path: "/dic_detail/list", Desc: "获取字典内容列表"},

		{Group: "系统服务", Method: "POST", Path: "/system/get_config", Desc: "获取配置文件内容"},
		{Group: "系统服务", Method: "POST", Path: "/system/set_config", Desc: "设置配置文件内容"},
		{Group: "系统服务", Method: "POST", Path: "/system/get_info", Desc: "获取服务器信息"},
		{Group: "系统服务", Method: "POST", Path: "/system/reload", Desc: "重启服务"},

		{Group: "操作记录", Method: "POST", Path: "/record/create", Desc: "新增操作记录"},
		{Group: "操作记录", Method: "DELETE", Path: "/record/delete", Desc: "删除操作记录"},
		{Group: "操作记录", Method: "DELETE", Path: "/record/deletes", Desc: "批量删除操作历史"},
		{Group: "操作记录", Method: "GET", Path: "/record/get", Desc: "根据ID获取操作记录"},
		{Group: "操作记录", Method: "GET", Path: "/record/list", Desc: "获取操作记录列表"},

		// {Group: "代码生成器", Method: "GET", Path: "/autoCode/getDB", Desc: "获取所有数据库"},
		// {Group: "代码生成器", Method: "GET", Path: "/autoCode/getTables", Desc: "获取数据库表"},
		// {Group: "代码生成器", Method: "POST", Path: "/autoCode/createTemp", Desc: "自动化代码"},
		// {Group: "代码生成器", Method: "POST", Path: "/autoCode/preview", Desc: "预览自动化代码"},
		// {Group: "代码生成器", Method: "GET", Path: "/autoCode/getColumn", Desc: "获取所选table的所有字段"},
		// {Group: "代码生成器", Method: "POST", Path: "/autoCode/createPlug", Desc: "自动创建插件包"},
		// {Group: "代码生成器", Method: "POST", Path: "/autoCode/installPlugin", Desc: "安装插件"},
		// {Group: "代码生成器", Method: "POST", Path: "/autoCode/pubPlug", Desc: "打包插件"},
		// {Group: "包生成器", Method: "POST", Path: "/autoCode/createPackage", Desc: "生成包(package)"},
		// {Group: "包生成器", Method: "POST", Path: "/autoCode/getPackage", Desc: "获取所有包(package)"},
		// {Group: "包生成器", Method: "POST", Path: "/autoCode/delPackage", Desc: "删除包(package)"},
		// {Group: "代码生成器历史", Method: "POST", Path: "/autoCode/getMeta", Desc: "获取meta信息"},
		// {Group: "代码生成器历史", Method: "POST", Path: "/autoCode/rollback", Desc: "回滚自动生成代码"},
		// {Group: "代码生成器历史", Method: "POST", Path: "/autoCode/getSysHistory", Desc: "查询回滚记录"},
		// {Group: "代码生成器历史", Method: "POST", Path: "/autoCode/delSysHistory", Desc: "删除回滚记录"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysApi{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initApi) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ? AND method = ?", "/jwt/jsonInBlacklist", "POST").
		First(&sysModel.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
