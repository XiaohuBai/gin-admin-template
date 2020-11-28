/*
 * @Author: XiaohuBai
 * @Date: 2020-11-25 10:00:25
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-11-26 00:16:34
 * @Description: 描述
 */
package router

import (
	v1 "gin-admin-template/api/v1"
	"gin-admin-template/middleware"

	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	{
		MenuRouter.POST("getMenu", v1.Login) // 获取菜单树
	}
	return MenuRouter
}
