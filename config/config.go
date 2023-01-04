package config

import (
	"github.com/bianjieai/opb-sdk-go/pkg/app/sdk/client"
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
		MaxConnectionNum   int    `mapstructure:"max_connection_num"`
		InitConnectionNum  int    `mapstructure:"init_connection_num"`
		GrpcAddress string `mapstructure:"grpcAddress"`
		ChainID string `mapstructure:"chainID"`
		Algo string `mapstructure:"algo"`
		Password string `mapstructure:"password"`
		Mnemonic string `mapstructure:"mnemonic"`
	}
)
const EnvNameConfigFilePath = "CONFIG_FILE_PATH"

func InitConfig()  (*Config,error) {
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
		return  nil,err
	}
	var cfg Config
	if err := rootViper.Unmarshal(&cfg); err != nil {
		return  nil,err
	}

	return &cfg, nil
}