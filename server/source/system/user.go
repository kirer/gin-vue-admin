package system

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	sysModel "kirer.cn/server/model/system"
	"kirer.cn/server/source"
	"kirer.cn/server/utils"
)

const initOrderUser = initOrderAuth + 1

type initUser struct{}

// auto run
func init() {
	source.RegisterInit(initOrderUser, &initUser{})
}

func (i *initUser) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, source.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysUser{})
}

func (i *initUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysUser{})
}

func (i initUser) InitializerName() string {
	return sysModel.SysUser{}.TableName()
}

func (i *initUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, source.ErrMissingDBContext
	}
	adminPassword := utils.BcryptHash("44021993")

	entities := []sysModel.SysUser{
		{
			UUID:      uuid.Must(uuid.NewV4()),
			Username:  "admin",
			Password:  adminPassword,
			NickName:  "Kirer",
			HeaderImg: "https://qmplusimg.henrongyi.top/gva_header.jpg",
			AuthId:    888,
			Phone:     "17611111111",
			Email:     "333333333@qq.com",
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysUser{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	authEntities, ok := ctx.Value(initAuth{}.InitializerName()).([]sysModel.SysAuth)
	if !ok {
		return next, errors.Wrap(source.ErrMissingDependentContext, "创建 [用户-权限] 关联失败, 未找到权限表初始化数据")
	}
	if err = db.Model(&entities[0]).Association("Auths").Replace(authEntities); err != nil {
		return next, err
	}
	return next, err
}

func (i *initUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record sysModel.SysUser
	if errors.Is(db.Where("username = ?", "admin").Preload("Auths").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return len(record.Auths) > 0 && record.Auths[0].AuthId == 888
}
