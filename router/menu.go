/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-07 22:58:23
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-10 17:17:19
 * @Description: 描述
 */

package router

import (
	v1 "gin-admin-template/api/v1"
	"gin-admin-template/middleware"

	"github.com/gin-gonic/gin"
)

// InitMenuRouter 初始化菜单路由组
func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	{
		MenuRouter.POST("roleMenus", v1.RoleMenus) //获取角色路由信息
		MenuRouter.POST("listMenus", v1.ListMenus) //获取菜单路由信息
	}
	return MenuRouter
}
