package response

import "gin-admin-template/model/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
