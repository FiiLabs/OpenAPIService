package api_router

import (
	"github.com/FiiLabs/OpenAPIService/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine )  {
	v1 := router.Group("/v1beta1")
	{
		v1.POST("/account", handler.AccountHandler)
		v1.POST("/test", handler.TestHandler)
	}
}