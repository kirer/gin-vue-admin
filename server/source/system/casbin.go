package system

import (
	"context"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"kirer.cn/server/source"
)

const initOrderCasbin = initOrderApi + 1

type initCasbin struct{}

// auto run
func init() {
	source.RegisterInit(initOrderCasbin, &initCasbin{})
}

func (i *initCasbin) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, source.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&adapter.CasbinRule{})
}

func (i *initCasbin) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&adapter.CasbinRule{})
}

func (i initCasbin) InitializerName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (i *initCasbin) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, source.ErrMissingDBContext
	}
	entities := []adapter.CasbinRule{
		{Ptype: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/admin_register", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/user/setUserInfo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/user/setSelfInfo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/deleteUser", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/setUserAuth", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/setUserAuthorities", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/resetPassword", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/api/create", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/delete", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/api/deletes", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/api/update", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/api/get", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/list", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/api/all", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/auth/create", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/auth/delete", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/auth/update", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/auth/list", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/auth/copy", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/auth/set_data", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/auth_btn/delete", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/auth_btn/update", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/auth_btn/get", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/casbin/update", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/casbin/get", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/menu/create", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/create_auth", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/delete", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/menu/update", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/menu/get", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/get_auth", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/get_current", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/list", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/all", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/dic/create", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/dic/delete", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/dic/update", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/dic/get", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/dic/list", V2: "GET"},

		{Ptype: "p", V0: "888", V1: "/dic_detail/create", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/dic_detail/delete", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/dic_detail/update", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/dic_detail/get", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/dic_detail/list", V2: "GET"},

		{Ptype: "p", V0: "888", V1: "/system/get_config", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/system/set_config", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/system/get_info", V2: "POST"},

		{Ptype: "p", V0: "888", V1: "/record/create", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/record/delete", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/record/deletes", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/record/get", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/record/list", V2: "GET"},

		// {Ptype: "p", V0: "888", V1: "/autoCode/getDB", V2: "GET"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/getMeta", V2: "POST"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/preview", V2: "POST"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/getTables", V2: "GET"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/getColumn", V2: "GET"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/rollback", V2: "POST"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/createTemp", V2: "POST"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/delSysHistory", V2: "POST"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/getSysHistory", V2: "POST"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/createPackage", V2: "POST"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/getPackage", V2: "POST"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/delPackage", V2: "POST"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/createPlug", V2: "POST"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/installPlugin", V2: "POST"},
		// {Ptype: "p", V0: "888", V1: "/autoCode/pubPlug", V2: "POST"},

	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, "Casbin 表 ("+i.InitializerName()+") 数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initCasbin) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where(adapter.CasbinRule{Ptype: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"}).
		First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
