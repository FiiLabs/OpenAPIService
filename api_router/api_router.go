package api_router

import (
	"github.com/FiiLabs/OpenAPIService/handler"
	"github.com/FiiLabs/OpenAPIService/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine )  {

	router.Use(middlewares.SignatureVerification())
	v1 := router.Group("/v1beta1")
	{
		v1.POST("/account", handler.AccountHandler)
		v1.POST("/test", handler.TestHandler)
	}
}