/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-02 17:30:18
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-07 22:59:37
 * @Description: 描述
 */

package v1

import (
	"gin-admin-template/global"
	"gin-admin-template/model/request"
	"gin-admin-template/model/response"
	"gin-admin-template/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RoleMenus 获取角色路由信息
func RoleMenus(c *gin.Context) {
	var r request.RoleReq
	_ = c.ShouldBindJSON(&r)

	if menus, err := service.RoleMenus(r.Role); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.MenusResponse{Menus: menus}, "获取成功", c)
	}
}

// ListMenus 获取菜单路由信息
func ListMenus(c *gin.Context) {
	if menus, err := service.ListMenus(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.MenusResponse{Menus: menus}, "获取成功", c)
	}
}

// getJwtData 获取jwt数据
func getJwtData(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户信息失败, 请检查路由是否使用jwt中间件")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.Role
	}
}
