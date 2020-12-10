/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-03 09:47:42
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-07 22:49:59
 * @Description: 描述
 */

package service

import (
	"gin-admin-template/global"
	"gin-admin-template/model"
)

// RoleMenus 获取角色路由信息
func RoleMenus(role string) (menus []model.Menu, err error) {
	roleMenuList, err := menusListByRole(role)
	m := make([]model.Menu, 0)
	for i := 0; i < len(roleMenuList); i++ {
		if roleMenuList[i].ParentID != 0 {
			continue
		}
		menusInfo := diguiMenu(&roleMenuList, roleMenuList[i])
		m = append(m, menusInfo)
	}
	return m, err
}

func menusListByRole(role string) (roleMenuList []model.Menu, err error) {
	err = global.GVA_DB.Table("menus").Select("menus.*").Joins("left join role_menus on menus.id = role_menus.menu_id where role_menus.role = ?", role).Scan(&roleMenuList).Error
	/* 	global.GVA_DB.Where("authority_id = ?", authorityId).Preload("Parameters").Find(&allMenus).Error
	   	err = global.GVA_DB.Model(model.Menu{}).Select("*").Joins("left join role_menus on menus.id = role_menus.menu_id where role_menus.role = ?", role).Scan(menus).Error */
	return
}

// ListMenus 获取菜单路由信息
func ListMenus() (MenuList []model.Menu, err error) {
	var menus = []model.Menu{}

	err = global.GVA_DB.Find(&menus).Error

	m := make([]model.Menu, 0)
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentID != 0 {
			continue
		}
		menusInfo := diguiMenu(&menus, menus[i])
		m = append(m, menusInfo)
	}
	return m, err
}

func diguiMenu(menulist *[]model.Menu, menu model.Menu) model.Menu {
	list := *menulist
	min := make([]model.Menu, 0)
	for j := 0; j < len(list); j++ {
		if menu.MenuID != list[j].ParentID {
			continue
		}
		mi := model.Menu{}
		mi.MenuID = list[j].MenuID
		mi.ParentID = list[j].ParentID
		mi.Path = list[j].Path
		mi.Name = list[j].Name
		mi.Path = list[j].Path
		mi.Hidden = list[j].Hidden
		mi.Component = list[j].Component
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Children = []model.Menu{}

		if mi.MenuType != "F" {
			ms := diguiMenu(menulist, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}
