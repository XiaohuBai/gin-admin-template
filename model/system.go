/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 15:48:24
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-01 15:48:36
 * @Description: 描述
 */

package model

import (
	"gin-admin-template/config"
)

// System 配置文件结构体
type System struct {
	Config config.Server
}
