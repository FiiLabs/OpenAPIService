package pool

import (
	"context"
	"errors"
	"fmt"
	"github.com/FiiLabs/OpenAPIService/config"
	opb "github.com/bianjieai/opb-sdk-go/pkg/app/sdk"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/store"
	commonPool "github.com/jolestar/go-commons-pool"
	"strings"
	"sync"
	"time"
)

var (
	poolObject  *commonPool.ObjectPool
	poolFactory PoolFactory
	ctx         = context.Background()
	_conf *config.Config
	_dao *store.KeyDAO
)

func Init(conf *config.Config,dao *store.KeyDAO) {
	var (
		syncMap sync.Map
	)
	_conf = conf
	_dao = dao
	nodeUrls := strings.Split(conf.Server.RpcAddress, ",")
	rpcUrls := strings.Split(conf.Server.GrpcAddress, ",")
	for index, url := range nodeUrls {
		key := generateId(url)
		endPoint := EndPoint{
			rpcAddress:   url,
			grpcAddress:   rpcUrls[index],
			Available: true,
		}

		syncMap.Store(key, endPoint)
	}

	poolFactory = PoolFactory{
		peersMap: syncMap,
	}

	config := commonPool.NewDefaultPoolConfig()
	config.MaxTotal = conf.Server.MaxConnectionNum
	config.MaxIdle = conf.Server.InitConnectionNum
	config.MinIdle = conf.Server.InitConnectionNum
	config.TestOnBorrow = true
	config.TestOnCreate = true
	config.TestWhileIdle = true

	poolObject = commonPool.NewObjectPool(ctx, &poolFactory, config)
	poolObject.PreparePool(ctx)
}
func GetConfig() *config.Config {
	return _conf
}
func newClient(nodeUrl string,grpcUrl string) (*Client, error) {
	fee, _ := types.ParseDecCoins("400000ugas")
	bech32AddressPrefix := types.AddrPrefixCfg{
		AccountAddr:   "metaos",
		ValidatorAddr: "metaosvaloper",
		ConsensusAddr: "metaosvalcons",
		AccountPub:    "metaospub",
		ValidatorPub:  "metaosvaloperpub",
		ConsensusPub:  "metaosvalconspub",
	}
	options := []types.Option{
		types.AlgoOption(_conf.Server.Algo),
		//types.KeyDAOOption(store.NewMemory(nil)),
		types.KeyDAOOption(*_dao),
		types.TimeoutOption(10),
		types.FeeOption(fee),
		types.CachedOption(true),
		types.Bech32AddressPrefixOption(&bech32AddressPrefix),
	}
	cfg, err := types.NewClientConfig(nodeUrl, grpcUrl, _conf.Server.ChainID, options...)
	if err != nil {
		fmt.Println(fmt.Errorf("new client error: %s", err.Error()))
		return nil,err
	}

	_client := opb.NewClient(cfg, nil)

	return &Client{
		Id:   generateId(nodeUrl),
		Client:&_client,
	}, err
}
// get client from pool
func GetClient() *Client {
	c, err := poolObject.BorrowObject(ctx)
	for err != nil {
		fmt.Println("GetClient failed,will try again after 3 seconds")
		time.Sleep(3 * time.Second)
		c, err = poolObject.BorrowObject(ctx)
	}

	return c.(*Client)
}

// release client
func (c *Client) Release() {
	err := poolObject.ReturnObject(ctx, c)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (c *Client) HeartBeat() error {
	http:= c.Client.BaseClient
	_, err := http.Health(context.Background())
	return err
}

func ClosePool() {
	poolObject.Close(ctx)
}

func GetClientWithTimeout(timeout time.Duration) (*Client, error) {
	c := make(chan interface{})
	errCh := make(chan error)
	go func() {
		client, err := poolObject.BorrowObject(ctx)
		if err != nil {
			errCh <- err
		} else {
			c <- client
		}
	}()
	select {
	case res := <-c:
		return res.(*Client), nil
	case res := <-errCh:
		return nil, res
	case <-time.After(timeout):
		return nil, errors.New("rpc node timeout")
	}
}