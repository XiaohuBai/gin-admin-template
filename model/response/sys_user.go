/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 14:50:22
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-01 14:51:45
 * @Description: 描述
 */
package response

import (
	"gin-admin-template/model"
)

type SysUserResponse struct {
	User model.User `json:"user"`
}

type LoginResponse struct {
	UserInfo  model.User `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}
