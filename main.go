/*
 * @Author: XiaohuBai
 * @Date: 2020-11-25 10:00:25
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-11-26 10:03:59
 * @Description: 程序入口文件，用于初始化各个模块
 */

package main

import (
	"gin-admin-template/core"
	"gin-admin-template/global"
	"gin-admin-template/initialize"
)

func main() {
	global.GVA_VP = core.Viper()          // 初始化Viper配置库：读取和监视配置文件
	global.GVA_LOG = core.Zap()           // 初始化zap日志库
	global.GVA_DB = initialize.Gorm()     // gorm连接数据库
	initialize.MysqlTables(global.GVA_DB) // 初始化表
	// 程序结束前关闭数据库链接
	db, _ := global.GVA_DB.DB()
	defer db.Close()

	core.RunServer()
}
