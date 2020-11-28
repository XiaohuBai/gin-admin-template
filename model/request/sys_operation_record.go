package request

import "gin-admin-template/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
