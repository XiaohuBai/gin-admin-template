/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 14:41:16
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-10 17:30:30
 * @Description:jwt要生成的token内容、配置
 */

package request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	UUID       uuid.UUID
	ID         uint
	Username   string
	Role       string
	Phone      string
	BufferTime int64
	jwt.StandardClaims
}
