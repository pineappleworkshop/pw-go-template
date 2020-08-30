package stores

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Hosts  []string `json:"hosts"`
	Client *mongo.Client
}

func NewMongo(hosts []string) *Mongo {
	return &Mongo{
		Hosts: hosts,
	}
}

func (m *Mongo) Connect() error {
	ctx := context.TODO()
	t := 30 * time.Second
	tr := true

	client, err := mongo.NewClient(&options.ClientOptions{
		ConnectTimeout: &t,
		Hosts:          m.Hosts,
		Direct:         &tr,
	})
	if err != nil {
		return err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return err
	}

	m.Client = client

	err = m.Client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
