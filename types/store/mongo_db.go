package store

import (
	"fmt"
	"github.com/FiiLabs/OpenAPIService/models/do"
	"github.com/irisnet/core-sdk-go/types/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	mdbDAO MongoDBDAO
)

type MongoDBDAO struct {
	db *do.MKeyInfo
	store.Crypto
}

// Use mongodb as storage
func NewMongoDB( crypto store.Crypto) (store.KeyDAO, error) {
	db := &do.MKeyInfo{}

	if crypto == nil {
		crypto = store.AES{}
	}

	mongoDB := MongoDBDAO{
		db:     db,
		Crypto: crypto,
	}
	mdbDAO.db = db
	return mongoDB, nil
}

func EnsureIndexes() {
	mdbDAO.db.EnsureIndexes()
}

func (k MongoDBDAO) Write(name, password string, info store.KeyInfo) error {
	var account do.MKeyInfo
	if k.Has(name) {
		return fmt.Errorf("name %s has exist", name)
	}

	privStr, err := k.Encrypt(info.PrivKeyArmor, password)
	if err != nil {
		return err
	}

	info.PrivKeyArmor = privStr

	k.CopyMemData(&info, &account)
	account.ID = primitive.NewObjectID()
	return k.db.InsertAccount(&account)
}

func (k MongoDBDAO) Read(name, password string) (store store.KeyInfo, err error) {
	bz, err := k.db.GetAccountByName(name)
	if bz == nil || err != nil {
		return store, err
	}

	k.CopyMongoData(bz, &store)

	if len(password) > 0 {
		privStr, err := k.Decrypt(store.PrivKeyArmor, password)
		if err != nil {
			return store, err
		}
		store.PrivKeyArmor = privStr
	}
	return
}
// ReadMetadata read a key information from the local store
func (k MongoDBDAO) ReadMetadata(name string) (store store.KeyInfo, err error) {
	bz, err := k.db.GetAccountByName(name)
	if bz == nil || err != nil {
		return store, err
	}

	k.CopyMongoData(bz, &store)
	return
}
func (k MongoDBDAO) Delete(name, password string) error {
	_, err := k.Read(name, password)
	if err != nil {
		return err
	}
	return k.db.DeleteAccount(name)
}

func (k MongoDBDAO) Has(name string) bool {
	existed, err := k.db.Exist(name)
	if err != nil {
		return false
	}
	return existed
}

func (k MongoDBDAO) CopyMongoData(src *do.MKeyInfo, dst *store.KeyInfo) {
	dst.Name = src.KeyName
	dst.Algo = src.Algo
	dst.PubKey = src.PubKey
	dst.PrivKeyArmor = src.PrivKeyArmor
}

func (k MongoDBDAO) CopyMemData(src *store.KeyInfo, dst *do.MKeyInfo) {
	dst.KeyName = src.Name
	dst.Algo = src.Algo
	dst.PubKey = src.PubKey
	dst.PrivKeyArmor = src.PrivKeyArmor
}