package stores

import (
	"{{<service_name>}}/config"
)

var DB *Store

func InitDbs() {
	if config.Conf.Env == config.WORKSTATION {
		DB = NewStores()
		if err := DB.Mongo.Connect(); err != nil {
			panic(err)
		}
	} else if config.Conf.Env == config.DEV {
		DB = NewStores()
		if err := DB.Mongo.Connect(); err != nil {
			panic(err)
		}
	} else if config.Conf.Env == config.STAGE {
		DB = NewStores()
		if err := DB.Mongo.Connect(); err != nil {
			panic(err)
		}
	}
}

type Store struct {
	Mongo *Mongo
}

func NewStores() *Store {
	if config.Conf.GetEnv() != config.WORKSTATION {
		return &Store{
			Mongo: NewMongo(MONGOHOSTS_CLUSTER),
		}
	}

	return &Store{
		Mongo: NewMongo(MONGOHOSTS_WORKSTATION),
	}
}