/*
 * @Author: XiaohuBai
 * @Date: 2020-11-25 10:00:25
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-02 17:43:56
 * @Description: 描述
 */

package router

import (
	v1 "gin-admin-template/api/v1"

	"github.com/gin-gonic/gin"
)

// InitBaseRouter 基础路由，不做鉴权
func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.GET("captcha", v1.Captcha)
		BaseRouter.GET("getSystemConfig", v1.GetSystemConfig)
		BaseRouter.GET("roleMenus", v1.roleMenus)
	}
	return BaseRouter
}
