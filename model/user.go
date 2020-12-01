/*
 * @Author: XiaohuBai
 * @Date: 2020-11-25 10:00:25
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-01 15:48:50
 * @Description: 描述
 */

package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//User 用户表结构
type User struct {
	gorm.Model
	UUID      uuid.UUID `gorm:"type:varchar(50);not null;default:'';comment:'uuid'" json:"uuid"`
	Role      string    `gorm:"type:varchar(1);not null;default:'0';comment:'角色 0:普通用户,1:管理员'" json:"role"`
	Username  string    `gorm:"type:varchar(50);not null;default:'';comment:'账户'" json:"username"`
	Password  string    `gorm:"type:varchar(50);not null;default:'';comment:'密码'" json:"password"`
	Nickname  string    `gorm:"type:varchar(50);not null;default:'昵称';comment:'昵称'" json:"nickname"`
	Phone     string    `gorm:"type:varchar(20);not null;default:'';comment:'手机号'" json:"phone"`
	Sex       string    `gorm:"type:varchar(1);not null;default:'0';comment:'性别 0:男,1:女'" json:"sex"`
	Email     string    `gorm:"type:varchar(50);not null;default:'';comment:'邮箱'" json:"email"`
	HeaderImg string    `gorm:"type:varchar(50);not null;default:'http://xiaohubai.cn/head.png';comment:'头像'" json:"headerImg"`
}
