/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-09 14:34:06
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-10 22:44:13
 * @Description: 描述
 */

package middleware

import (
	"gin-admin-template/global"
	"gin-admin-template/model/request"
	"gin-admin-template/model/response"
	"gin-admin-template/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CasbinHandler 权限拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色ID
		sub := strconv.FormatUint(uint64(waitUse.ID), 10)
		e := service.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.GVA_CONFIG.System.Env == "develop" || success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
