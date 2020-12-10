/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-09 14:50:51
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-09 15:10:44
 * @Description: 描述
 */

package service

import (
	"errors"
	"gin-admin-template/global"
	"gin-admin-template/model"
	"gin-admin-template/model/request"
	"strings"

	"github.com/casbin/casbin/util"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

// UpdateCasbin 更新casbin权限
func UpdateCasbin(RoleID string, casbinInfos []request.CasbinInfo) error {
	ClearCasbin(0, RoleID)
	rules := [][]string{}
	for _, v := range casbinInfos {
		cm := model.CasbinModel{
			Ptype:  "p",
			RoleID: RoleID,
			Path:   v.Path,
			Method: v.Method,
		}
		rules = append(rules, []string{cm.RoleID, cm.Path, cm.Method})
	}
	e := Casbin()
	success, _ := e.AddPolicies(rules)
	if success == false {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

// UpdateCasbinAPI API更新随动
func UpdateCasbinAPI(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.GVA_DB.Table("casbin_rule").Model(&model.CasbinModel{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// GetPolicyPathByAuthorityID 获取权限列表
func GetPolicyPathByAuthorityID(RoleID string) (pathMaps []request.CasbinInfo) {
	e := Casbin()
	list := e.GetFilteredPolicy(0, RoleID)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// ClearCasbin 清除匹配的权限
func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success

}

//Casbin 持久化到数据库  引入自定义规则
func Casbin() *casbin.Enforcer {
	admin := global.GVA_CONFIG.Mysql
	a, _ := gormadapter.NewAdapter(global.GVA_CONFIG.System.DbType, admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname, true)
	e, _ := casbin.NewEnforcer(global.GVA_CONFIG.Casbin.ModelPath, a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}

// ParamsMatch 自定义规则函数
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// ParamsMatchFunc 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}
