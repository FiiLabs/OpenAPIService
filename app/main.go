package main

import (
	"fmt"
	"github.com/FiiLabs/OpenAPIService/api_router"
	"github.com/FiiLabs/OpenAPIService/config"
	"github.com/FiiLabs/OpenAPIService/libs/pool"
	"github.com/FiiLabs/OpenAPIService/models/do"
	"github.com/FiiLabs/OpenAPIService/types/store"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	conf, err := config.InitConfig()
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
	do.Init(conf)
	store.EnsureIndexes()
	dao,_ := store.NewMongoDB(nil)
	pool.Init(conf,&dao)
	router := gin.Default()

	api_router.RegisterRouter(router)

	router.Run(":30000")
}