package main

import (
	"fmt"
	"github.com/FiiLabs/OpenAPIService/api_router"
	"github.com/FiiLabs/OpenAPIService/config"
	"github.com/FiiLabs/OpenAPIService/models/do"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		fmt.Println("System Exit")

		do.Close()

		if err := recover(); err != nil {
			os.Exit(1)
		}
	}()
	do.Init()
	router := gin.Default()

	api_router.RegisterRouter(router)

	router.Run(":8080")
}