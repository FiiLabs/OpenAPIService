package api_router

import (
	"github.com/FiiLabs/OpenAPIService/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine )  {

	//router.Use(middlewares.SignatureVerification())
	v1 := router.Group("/v1beta1")
	{
		v1.POST("/account", handler.AccountHandler)
		v1.POST("/test", handler.TestHandler)
		v1.POST("/nft/classes", handler.NFTClassHandler)
		v1.POST("/nft/class-transfers", handler.NFTClassTransferHandler)
		v1.POST("/nft/nfts", handler.NFTHandler)
		v1.POST("/nft/nft-transfers", handler.NFTTransferHandler)
		v1.PATCH("/nft/nfts", handler.NFTEditHandler)
	}
}