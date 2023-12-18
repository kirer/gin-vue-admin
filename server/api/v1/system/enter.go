package system

import "kirer.cn/server/service"

type ApiGroup struct {
	JwtApi
	BaseApi
	SystemApi
	CasbinApi
	AutoCodeApi
	SystemApiApi
	AuthApi
	DicApi
	AuthMenuApi
	RecordApi
	AutoCodeHistoryApi
	DicDetailApi
	AuthBtnApi
}

var (
	apiService              = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService              = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService             = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
	casbinService           = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	autoCodeService         = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
	baseMenuService         = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	authService             = service.ServiceGroupApp.SystemServiceGroup.AuthService
	dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DicService
	systemConfigService     = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	operationRecordService  = service.ServiceGroupApp.SystemServiceGroup.RecordService
	autoCodeHistoryService  = service.ServiceGroupApp.SystemServiceGroup.AutoCodeHistoryService
	dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DicDetailService
	authBtnService          = service.ServiceGroupApp.SystemServiceGroup.AuthBtnService
)
