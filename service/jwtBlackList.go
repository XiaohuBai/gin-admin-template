/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 14:53:48
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-09 14:19:53
 * @Description: 描述
 */

package service

import (
	"errors"
	"gin-admin-template/global"
	"gin-admin-template/model"
	"time"

	"gorm.io/gorm"
)

// JSONInBlacklist 作废token 插入数据库
func JSONInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	return
}

//IsBlacklist 查询是否是黑名单token
func IsBlacklist(jwt string) bool {
	isNotFound := errors.Is(global.GVA_DB.Where("jwt = ?", jwt).First(&model.JwtBlacklist{}).Error, gorm.ErrRecordNotFound)
	return !isNotFound
}

// GetRedisJWT 从redis获取token
func GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(userName).Result()
	return redisJWT, err
}

// SetRedisJWT 从redis设置token
func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := 60 * 60 * 24 * 7 * time.Second
	err = global.GVA_REDIS.Set(userName, jwt, timer).Err()
	return err
}
