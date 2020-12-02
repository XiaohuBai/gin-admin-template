/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 23:16:43
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-01 23:43:52
 * @Description: 描述
 */

package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5V md5加密
func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
