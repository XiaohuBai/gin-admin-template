/*
 * @Author: XiaohuBai
 * @Date: 2020-11-25 10:00:25
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-11-25 23:54:37
 * @Description: 描述
 */
package model

import (
	"gin-admin-template/global"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username    string    `json:"userName" gorm:"comment:用户登录名"`
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`
	NickName    string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
	HeaderImg   string    `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	AuthorityId string    `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}
