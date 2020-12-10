/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 15:39:01
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-07 22:50:13
 * @Description: 描述
 */

package service

import (
	"gin-admin-template/config"
	"gin-admin-template/global"
	"gin-admin-template/model"
	"gin-admin-template/utils"

	"go.uber.org/zap"
)

//GetSystemConfig 读取配置文件
func GetSystemConfig() (conf config.Server, err error) {
	return global.GVA_CONFIG, nil
}

// SetSystemConfig 设置配置文件
func SetSystemConfig(system model.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}

//GetServerInfo 获取服务器信息
func GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.CPU, err = utils.InitCPU(); err != nil {
		global.GVA_LOG.Error("func utils.InitCPU() Failed!", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		global.GVA_LOG.Error("func utils.InitRAM() Failed!", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.GVA_LOG.Error("func utils.InitDisk() Failed!", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
