/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-02 17:30:18
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-02 18:23:15
 * @Description: 描述
 */

package v1

import (
	"gin-admin-template/global"
	"gin-admin-template/model/request"
	"gin-admin-template/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func roleMenus(c *gin.Context) {
	if err, menus := service.roleMenus(getJwtData(c)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// TODO
func getJwtData(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.Role
	}
}
