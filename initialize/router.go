/*
 * @Author: XiaohuBai
 * @Date: 2020-11-25 10:00:25
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-01 13:54:49
 * @Description: 初始化总路由
 */
package initialize

import (
	_ "gin-admin-template/docs"
	"gin-admin-template/global"
	"gin-admin-template/middleware"
	"gin-admin-template/router"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Routers 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	//Router.Use(middleware.LoadTLS())                                                   	// 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	/* 	PrivateGroup := Router.Group("")
	   	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	   	{
	   		router.InitMenuRouter(PrivateGroup) // 注册menu路由
	   	} */
	global.GVA_LOG.Info("router register success")
	return Router
}
