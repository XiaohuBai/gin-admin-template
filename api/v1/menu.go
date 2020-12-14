/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-02 17:30:18
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-11 17:58:23
 * @Description: 描述
 */

package v1

import (
	"gin-admin-template/global"
	"gin-admin-template/model/response"
	"gin-admin-template/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RoleMenus 获取角色路由信息
func RoleMenus(c *gin.Context) {
	waitUse := getJwtData(c)
	if menus, err := service.RoleMenus(strconv.FormatUint(uint64(waitUse.ID), 10)); err != nil {
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
