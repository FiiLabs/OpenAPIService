package config

import (
	"fmt"
	opb "github.com/bianjieai/opb-sdk-go/pkg/app/sdk"
	"github.com/bianjieai/opb-sdk-go/pkg/app/sdk/client"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/store"
)
var (

	rpcAddress  = "http://localhost:26657"
	grpcAddress = "localhost:9090"
	chainID     = "testnet"
	algo        = "sm2"
)

type (
	config struct {
		Client client.Client
	}
)

var Conf config

func InitConfig()  {
	fee, _ := types.ParseDecCoins("400000ugas")
	options := []types.Option{
		types.AlgoOption(algo),
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(10),
		types.FeeOption(fee),
		types.CachedOption(true),
	}
	cfg, err := types.NewClientConfig(rpcAddress, grpcAddress, chainID, options...)
	if err != nil {
		fmt.Println(fmt.Errorf("new client error: %s", err.Error()))
		return
	}

	Conf.Client = opb.NewClient(cfg, nil)
}

func GetConfig() *config {
	return &Conf
}

func GetConfigClient() *client.Client {
	return &Conf.Client
}