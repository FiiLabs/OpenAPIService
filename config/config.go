package config

import (
	"fmt"
	opb "github.com/bianjieai/opb-sdk-go/pkg/app/sdk"
	"github.com/bianjieai/opb-sdk-go/pkg/app/sdk/client"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/store"
	"github.com/spf13/viper"
	"os"
)

type (
	Config struct {
		DataBase DataBaseConf `mapstructure:"database"`
		Server   ServerConf   `mapstructure:"server"`
		Client client.Client
	}
	DataBaseConf struct {
		NodeUri  string `mapstructure:"node_uri"`
		Database string `mapstructure:"database"`
	}
	ServerConf struct {
		RpcAddress   string `mapstructure:"rpcAddress"`
		GrpcAddress string `mapstructure:"grpcAddress"`
		ChainID string `mapstructure:"chainID"`
		Algo string `mapstructure:"algo"`
		Password string `mapstructure:"password"`
		Mnemonic string `mapstructure:"mnemonic"`
	}
)
const EnvNameConfigFilePath = "CONFIG_FILE_PATH"

var conf Config

func InitConfig()  error {
	var ConfigFilePath string

	websit, found := os.LookupEnv(EnvNameConfigFilePath)
	if found {
		ConfigFilePath = websit
	} else {
		panic("not found CONFIG_FILE_PATH")
	}

	rootViper := viper.New()
	// Find home directory.
	rootViper.SetConfigFile(ConfigFilePath)

	// Find and read the config file
	if err := rootViper.ReadInConfig(); err != nil { // Handle errors reading the config file
		return  err
	}

	if err := rootViper.Unmarshal(&conf); err != nil {
		return  err
	}

	fee, _ := types.ParseDecCoins("400000ugas")
	options := []types.Option{
		types.AlgoOption(conf.Server.Algo),
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(10),
		types.FeeOption(fee),
		types.CachedOption(true),
	}
	cfg, err := types.NewClientConfig(conf.Server.RpcAddress, conf.Server.GrpcAddress, conf.Server.ChainID, options...)
	if err != nil {
		fmt.Println(fmt.Errorf("new client error: %s", err.Error()))
		return err
	}

	conf.Client = opb.NewClient(cfg, nil)
	return nil
}

func GetConfig() *Config {
	return &conf
}

func GetConfigClient() *client.Client {
	return &conf.Client
}