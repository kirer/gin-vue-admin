package response

import "kirer.cn/server/model/system"

type SysAuthResponse struct {
	Auth system.SysAuth `json:"auth"`
}

type SysAuthCopyResponse struct {
	Auth      system.SysAuth `json:"auth"`
	OldAuthId uint           `json:"oldAuthId"` // 旧角色ID
}
