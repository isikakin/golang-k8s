package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoClient interface {
	NewSession(dbName string) *mongo.Database
	Ping(ctx context.Context) error
}

type mongoClient struct {
	client *mongo.Client
}

func (self *mongoClient) NewSession(dbName string) *mongo.Database {
	return self.client.Database(dbName)
}

func (self *mongoClient) Ping(ctx context.Context) error {
	return self.client.Ping(ctx, nil)
}

func NewClient(connectionString, replicaSet string, timeout time.Duration) (MongoClient, error) {

	var (
		client *mongo.Client
		err    error
	)

	clientOptions := options.
		Client().
		ApplyURI(connectionString).SetReplicaSet(replicaSet)

	ctxWithTimeout, _ := context.WithTimeout(context.Background(), timeout)

	if client, err = mongo.Connect(ctxWithTimeout, clientOptions); err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &mongoClient{client: client}, nil
}
