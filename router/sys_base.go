/*
 * @Author: XiaohuBai
 * @Date: 2020-11-25 10:00:25
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-11-25 15:33:30
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
		BaseRouter.POST("captcha", v1.Captcha)
	}
	return BaseRouter
}
