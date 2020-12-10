/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-09 15:43:19
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-10 17:53:21
 * @Description: 描述
 */
package service

import (
	"gin-admin-template/global"
	"gin-admin-template/model"
)

// CreateSysOperationRecord 创建记录
func CreateSysOperationRecord(sysOperationRecord model.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}
