/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-09 15:05:22
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-09 15:05:41
 * @Description: 描述
 */

package model

type CasbinModel struct {
	Ptype  string `json:"ptype" gorm:"column:p_type"`
	RoleID string `json:"roleID" gorm:"column:v0"`
	Path   string `json:"path" gorm:"column:v1"`
	Method string `json:"method" gorm:"column:v2"`
}
