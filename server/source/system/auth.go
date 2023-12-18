package system

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	sysModel "kirer.cn/server/model/system"
	"kirer.cn/server/source"
	"kirer.cn/server/utils"
)

const initOrderAuth = initOrderCasbin + 1

type initAuth struct{}

// auto run
func init() {
	source.RegisterInit(initOrderAuth, &initAuth{})
}

func (i *initAuth) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, source.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysAuth{})
}

func (i *initAuth) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysAuth{})
}

func (i initAuth) InitializerName() string {
	return sysModel.SysAuth{}.TableName()
}

func (i *initAuth) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, source.ErrMissingDBContext
	}
	entities := []sysModel.SysAuth{
		{AuthId: 888, AuthName: "管理员", ParentId: utils.Pointer[uint](0), DefaultRouter: "dashboard"},
	}

	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败!", sysModel.SysAuth{}.TableName())
	}
	// data auth
	if err := db.Model(&entities[0]).Association("DataAuthId").Replace([]*sysModel.SysAuth{{AuthId: 888}}); err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败!", db.Model(&entities[0]).Association("DataAuthId").Relationship.JoinTable.Name)
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initAuth) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("auth_id = ?", "888").First(&sysModel.SysAuth{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
