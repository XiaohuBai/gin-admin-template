/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 15:13:27
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-01 15:42:09
 * @Description: 描述
 */

package v1

import (
	"gin-admin-template/global"
	"gin-admin-template/model"
	"gin-admin-template/model/response"
	"gin-admin-template/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//JSONInBlacklist token加入黑名单
func JSONInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := model.JwtBlacklist{Jwt: token}
	if err := service.JSONInBlacklist(jwt); err != nil {
		global.GVA_LOG.Error("jwt作废失败!", zap.Any("err", err))
		response.FailWithMessage("jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
