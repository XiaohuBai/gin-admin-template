/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 23:15:44
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-01 23:16:30
 * @Description: 描述
 */

package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// StructToMap 利用反射将结构体转化为map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

// ArrayToString 将数组格式化为字符串
func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
