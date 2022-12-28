package do

import (
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	CollectionNameBlock = "apikey"
)

type (
	APIKey struct {
		ApiKey     string `bson:"apikey"`
		ApiSecret   string  `bson:"apisecret"`
	}
)

func (d *APIKey) Name() string {
		return CollectionNameBlock
}

func (d *APIKey) GetSecret(key string) (string, error) {
	var result APIKey

	getApiSecret := func(c *qmgo.Collection) error {
		return c.Find(_ctx, bson.M{"apikey":key}).Limit(1).One(&result)
	}

	if err := ExecCollection(d.Name(), getApiSecret); err != nil {
		return "", err
	}

	return result.ApiSecret, nil
}