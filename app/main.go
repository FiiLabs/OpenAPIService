package main

import (
	"github.com/FiiLabs/OpenAPIService/api_router"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api_router.RegisterRouter(router)

	router.Run(":8080")
}