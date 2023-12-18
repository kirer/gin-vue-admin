package request

import (
	"kirer.cn/server/model/common/request"
	"kirer.cn/server/model/system"
)

type SysDicDetailSearch struct {
	system.SysDicDetail
	request.PageInfo
}
