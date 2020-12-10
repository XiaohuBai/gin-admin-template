/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-03 00:19:10
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-03 09:55:40
 * @Description: 描述 角色路由表
 */

package response

import "gin-admin-template/model"

//MenusResponse 返回的菜单路由
type MenusResponse struct {
	Menus []model.Menu `json:"menus"`
}
