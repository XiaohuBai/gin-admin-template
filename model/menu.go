/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-03 00:19:10
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-03 17:55:46
 * @Description: 描述 角色路由表
 */

package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	MenuID    int    `json:"menuId" gorm:"type:int(11)"`
	ParentID  int    `json:"parentId" gorm:"comment:父菜单ID"`
	MenuType  string `json:"menuType" gorm:"type:varchar(1);"`
	Path      string `json:"path" gorm:"comment:路由path"`
	Name      string `json:"name" gorm:"comment:路由name"`
	Hidden    int    `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component string `json:"component" gorm:"comment:对应前端文件路径"`
	Title     string `json:"title" gorm:"comment:菜单名"`
	Icon      string `json:"icon" gorm:"comment:菜单图标"`
	Children  []Menu `json:"children" gorm:"-"`
}

type RoleMenu struct {
	gorm.Model
	MenuID int    `json:"menuId" gorm:"type:int(11)"`
	Role   string `json:"role" gorm:"type:varchar(128)"`
}
