package do

import (
	"context"
	"fmt"
	"github.com/FiiLabs/OpenAPIService/config"
	"github.com/qiniu/qmgo"
)
var (
	_ctx  = context.Background()
	_conf *config.Config
	_cli  *qmgo.Client
)
func Init() {
	_conf = config.GetConfig()
	var maxPoolSize uint64 = 4096

	client, err := qmgo.NewClient(_ctx, &qmgo.Config{
		Uri:         _conf.DataBase.NodeUri,
		Database:    _conf.DataBase.Database,
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
func ExecCollection(collectionName string, s func(*qmgo.Collection) error) error {
	c := _cli.Database(GetDbConf().Database).Collection(collectionName)
	return s(c)
}