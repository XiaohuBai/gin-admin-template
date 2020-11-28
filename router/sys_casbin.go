package router

import (
	"gin-admin-template/api/v1"
	"gin-admin-template/middleware"
	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	CasbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	{
		CasbinRouter.POST("updateCasbin", v1.UpdateCasbin)
		CasbinRouter.POST("getPolicyPathByAuthorityId", v1.GetPolicyPathByAuthorityId)
	}
}
