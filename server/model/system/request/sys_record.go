package request

import (
	"kirer.cn/server/model/common/request"
	"kirer.cn/server/model/system"
)

type SysRecordSearch struct {
	system.SysRecord
	request.PageInfo
}
