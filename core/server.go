/*
 * @Author: XiaohuBai
 * @Date: 2020-11-25 10:00:25
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-01 16:15:45
 * @Description: 描述
 */

package core

import (
	"fmt"
	"gin-admin-template/global"
	"gin-admin-template/initialize"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

// RunServer 运行服务
func RunServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 gin-admin-template
	当前版本:V2.3.7
	默认自动化文档地址:http://localhost%s/swagger/index.html
	默认前端文件运行地址:http://localhost:8080
	如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-admin-template.com/docs/coffee
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
