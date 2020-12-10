/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 15:45:36
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-01 15:47:51
 * @Description: 描述
 */

package model

import "gorm.io/gorm"

//JwtBlacklist jwt结构体
type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
