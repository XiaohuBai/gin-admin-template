/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-02 14:05:06
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-02 14:08:19
 * @Description: 描述
 */

package response

// SysConfigResponse 系统基本信息
type SysConfigResponse struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
}
