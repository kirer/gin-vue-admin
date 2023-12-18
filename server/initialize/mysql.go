package initialize

import (
	"database/sql"
	"fmt"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"kirer.cn/server/config"
	"kirer.cn/server/global"
	"kirer.cn/server/initialize/internal"
	"kirer.cn/server/model/system"
	"kirer.cn/server/source"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	m := global.CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	db, err := ensureDB(m)
	if err != nil {
		panic(fmt.Errorf(">>数据库连接失败:%s", err))
	}
	fmt.Printf(">>数据库连接成功:%s\n", m.Dsn())
	source.EnsureTableData(db)
	db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	err = registerTables(db) // 初始化表
	if err != nil {
		panic(fmt.Errorf(">>数据库注册表失败:%s", err))
	}
	fmt.Println(">>数据库初始化表成功")
	return db
}
func ensureDB(m config.Mysql) (db *gorm.DB, err error) {
	config := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	db, err = gorm.Open(mysql.New(config), internal.Gorm.Config(m.Prefix, m.Singular))
	if err == nil {
		return
	}
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", m.Dbname)
	if err = createDatabase(m.EmptyDsn(), "mysql", createSql); err != nil {
		return
	}
	fmt.Printf(">>数据库创建成功:%s\n", m.EmptyDsn())
	db, err = gorm.Open(mysql.New(config), internal.Gorm.Config(m.Prefix, m.Singular))
	if err == nil {
		return
	}
	global.CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return
}
func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
func registerTables(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		// 系统模块表
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuth{},
		system.SysDic{},
		system.SysRecord{},
		system.SysAutoCodeHistory{},
		system.SysDicDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthBtn{},
		system.SysAutoCode{},
	)
	return
}
