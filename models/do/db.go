package do

import (
	"context"
	"fmt"
	"github.com/FiiLabs/OpenAPIService/config"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
)
var (
	_ctx  = context.Background()
	_conf *config.Config
	_cli  *qmgo.Client
)
func Init(conf *config.Config) {
	var maxPoolSize uint64 = 4096
	_conf = conf
	client, err := qmgo.NewClient(_ctx, &qmgo.Config{
		Uri:         conf.DataBase.NodeUri,
		Database:    conf.DataBase.Database,
		MaxPoolSize: &maxPoolSize,
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("connect mongo failed, uri: %s, err:%s", _conf.DataBase.NodeUri, err.Error()))
	}
	_cli = client

	fmt.Println("init db success")

	return
}
func GetClient() *qmgo.Client {
	return _cli
}
func Close() {
	fmt.Println("release resource :mongoDb")
	if _cli != nil {
		_cli.Close(_ctx)
	}
}
func GetDbConf() *config.DataBaseConf {
	if _conf == nil {
		panic("db.Init not work")
	}
	return &_conf.DataBase
}
func GetSrvConf() *config.ServerConf {
	if _conf == nil {
		panic("db.Init not work")
	}
	return &_conf.Server
}

func ensureIndexes(collectionName string, indexes []options.IndexModel) {
	c := _cli.Database(GetDbConf().Database).Collection(collectionName)
	if len(indexes) > 0 {
		for _, v := range indexes {
			if err := c.CreateOneIndex(context.Background(), v); err != nil {
				fmt.Println("ensure index fail")
			}
		}
	}
}

func ExecCollection(collectionName string, s func(*qmgo.Collection) error) error {
	c := _cli.Database(GetDbConf().Database).Collection(collectionName)
	return s(c)
}