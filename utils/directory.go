/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 19:24:26
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-01 19:24:49
 * @Description: 描述
 */

package utils

import (
	"gin-admin-template/global"
	"os"

	"go.uber.org/zap"
)

// PathExists 文件目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 批量创建文件夹
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.GVA_LOG.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.GVA_LOG.Error("create directory"+v, zap.Any(" error:", err))
			}
		}
	}
	return err
}
