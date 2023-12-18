package system

import (
	"context"
	"fmt"
	"sort"

	"github.com/gofrs/uuid/v5"
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"kirer.cn/server/config"
	"kirer.cn/server/global"
	"kirer.cn/server/utils"
)

// orderedInitializer 组合一个顺序字段，以供排序
type orderedInitializer struct {
	order int
	SubInitializer
}
type SubInitializer interface {
	InitializerName() string
	MigrateTable(ctx context.Context) (next context.Context, err error)
	InitializeData(ctx context.Context) (next context.Context, err error)
	TableCreated(ctx context.Context) bool
	DataInserted(ctx context.Context) bool
}

// initSlice 供 initializer 排序依赖时使用
type initSlice []*orderedInitializer

func (a initSlice) Len() int {
	return len(a)
}

func (a initSlice) Less(i, j int) bool {
	return a[i].order < a[j].order
}

func (a initSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

var (
	initializers initSlice
	cache        map[string]*orderedInitializer
)

func ensureTableData() (err error) {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", len(initializers))
	ctx := context.TODO()
	if len(initializers) == 0 {
		return
	}
	sort.Sort(&initializers) // 保证有依赖的 initializer 排在后面执行
	// Note: 若 initializer 只有单一依赖，可以写为 B=A+1, C=A+1; 由于 BC 之间没有依赖关系，所以谁先谁后并不影响初始化
	// 若存在多个依赖，可以写为 C=A+B, D=A+B+C, E=A+1;
	// C必然>A|B，因此在AB之后执行，D必然>A|B|C，因此在ABC后执行，而E只依赖A，顺序与CD无关，因此E与CD哪个先执行并不影响
	if err = initTables(ctx, initializers); err != nil {
		return err
	}
	if err = initData(ctx, initializers); err != nil {
		return err
	}
	if err = writeConfig(ctx); err != nil {
		return err
	}
	initializers = initSlice{}
	cache = map[string]*orderedInitializer{}
	return
}
func initTables(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.TableCreated(next) {
			continue
		}
		if n, err := init.MigrateTable(next); err != nil {
			return err
		} else {
			next = n
		}
	}
	return nil
}

func initData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.DataInserted(next) {
			color.Info.Printf("\n>>数据库>%v>初始数据已存在!\n", init.InitializerName())
			continue
		}
		if n, err := init.InitializeData(next); err != nil {
			color.Info.Printf("\n>>数据库>%v>初始数据失败!>%+v\n", init.InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf("\n>>数据库>%v>初始数据成功!\n", init.InitializerName())
		}
	}
	color.Info.Printf("\n>>数据库>初始数据成功!\n")
	return nil
}

func writeConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Mysql)
	if !ok {
		return errors.New("mysql config invalid")
	}
	global.CONFIG.Mysql = c
	global.CONFIG.JWT.SigningKey = uuid.Must(uuid.NewV4()).String()
	cs := utils.StructToMap(global.CONFIG)
	for k, v := range cs {
		global.VP.Set(k, v)
	}
	return global.VP.WriteConfig()
}
