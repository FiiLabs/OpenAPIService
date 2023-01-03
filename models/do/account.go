package do

import (
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionNameAccount = "account"
)

// KeyInfo saves the basic information of the key
type MKeyInfo struct {
	ID             primitive.ObjectID `bson:"_id"`
	KeyName         string `json:"name" bson:"name"`
	PubKey       []byte `json:"pubkey" bson:"pubkey"`
	PrivKeyArmor string `json:"priv_key_armor" bson:"priv_key_armor"`
	Algo         string `json:"algo" bson:"algo"`
}

func (d *MKeyInfo) Name() string {
	return CollectionNameAccount
}
func (d *MKeyInfo) EnsureIndexes() {
	var indexes []options.IndexModel
	indexes = append(indexes, options.IndexModel{
		Key:        []string{"-name"},
		Unique:     true,
		Background: true,
	})
	ensureIndexes(d.Name(), indexes)
}
func (d *MKeyInfo) Exist(name string) (bool, error) {
	var result MKeyInfo

	hasAccount := func(c *qmgo.Collection) error {
		return c.Find(_ctx, bson.M{"name":name}).Limit(1).One(&result)
	}

	if err := ExecCollection(d.Name(), hasAccount); err != nil {
		return false, err
	}

	return true, nil
}

func (d *MKeyInfo) GetAccountByName(name string) (*MKeyInfo, error) {
	var result MKeyInfo

	hasAccount := func(c *qmgo.Collection) error {
		return c.Find(_ctx, bson.M{"name":name}).Limit(1).One(&result)
	}

	if err := ExecCollection(d.Name(), hasAccount); err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *MKeyInfo) InsertAccount(account *MKeyInfo) error {
	fn := func(c *qmgo.Collection) error {
		_, err := c.InsertOne(_ctx, account)
		return err
	}
	return ExecCollection(d.Name(), fn)
}

func (d *MKeyInfo) DeleteAccount(name string) error {
	keyInfo,err := d.GetAccountByName(name)
	if err != nil {
		return err
	}

	fn := func(c *qmgo.Collection) error {
		return c.RemoveId(_ctx, keyInfo.ID)
	}
	return ExecCollection(d.Name(), fn)
}