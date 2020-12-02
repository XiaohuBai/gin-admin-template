/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 15:40:12
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-02 15:32:40
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

//GetSystemConfig 获取配置文件内容
func GetSystemConfig(c *gin.Context) {
	if config, err := service.GetSystemConfig(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysConfigResponse{
			Name: config.System.Name,
			Logo: config.System.Logo,
		}, "获取成功", c)
	}
}

//SetSystemConfig 置配置文件内容
func SetSystemConfig(c *gin.Context) {
	var sys model.System
	_ = c.ShouldBindJSON(&sys)
	if err := service.SetSystemConfig(sys); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Any("err", err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithData("设置成功", c)
	}
}
