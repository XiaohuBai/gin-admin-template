package response

import "gin-admin-template/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
