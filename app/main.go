package main

import (
	"github.com/FiiLabs/OpenAPIService/api_router"
	"github.com/FiiLabs/OpenAPIService/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	router := gin.Default()

	api_router.RegisterRouter(router)

	router.Run(":8080")
}