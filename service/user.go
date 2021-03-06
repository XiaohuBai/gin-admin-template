/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 15:15:01
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-09 14:12:39
 * @Description: 描述
 */

package service

import (
	"errors"
	"gin-admin-template/global"
	"gin-admin-template/model"
	"gin-admin-template/utils"
)

// Login 数据库查询用户信息
func Login(u *model.User) (userInter *model.User, err error) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return &user, err
}

//ChangePassword 修改密码
func ChangePassword(u *model.User, newPassword string) (userInter *model.User, err error) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return u, err
}

// FindUserByUUID 根据uuid查找用户
func FindUserByUUID(uuid string) (user *model.User, err error) {
	var u model.User
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}
