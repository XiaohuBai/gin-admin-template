/*
 * @Author: XiaohuBai
 * @Date: 2020-11-26 00:01:54
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-11-26 00:02:03
 * @Description: 描述
 */

package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string
}

type Empty struct{}
