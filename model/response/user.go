/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 14:50:22
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-02 15:41:54
 * @Description: 描述
 */

package response

import (
	"gin-admin-template/model"

	uuid "github.com/satori/go.uuid"
)

type SysUserResponse struct {
	User model.User `json:"user"`
}

type UserInfo struct {
	Username  string    `json:"username"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Nickname  string    `json:"nickname"`
	UUID      uuid.UUID `json:"uuid"`
	HeaderImg string    `json:"headerImg"`
	Sex       string    `json:"sex"`
}
type LoginResponse struct {
	UserInfo  UserInfo `json:"userInfo"`
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expiresAt"`
}
