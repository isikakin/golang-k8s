package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-k8s/internal/domain/enitity"
	mongoclient "golang-k8s/pkg/mongodb"
	"time"
)

const productCollectionName = "product"

type ProductRepository interface {
	GetAll(ctx context.Context) ([]enitity.Product, error)
	GetById(ctx context.Context, id string) (*enitity.Product, error)
	Insert(ctx context.Context, product *enitity.Product) error
	UpdateById(ctx context.Context, id string, product *enitity.Product) error
	RemoveById(ctx context.Context, id string) error
	Ping(ctx context.Context) error
}

type productRepository struct {
	mongoClient  mongoclient.MongoClient
	databaseName string
}

func (self *productRepository) Ping(ctx context.Context) error {
	return self.mongoClient.Ping(ctx)
}

func (self *productRepository) GetAll(ctx context.Context) ([]enitity.Product, error) {

	var session = self.mongoClient.NewSession(self.databaseName)

	var result []enitity.Product

	cur, err := session.Collection(productCollectionName).Find(ctx, bson.D{})

	if err != nil {
		if err == mongo.ErrNilDocument {
			return nil, nil
		}
		return nil, err
	}

	defer cur.Close(ctx)

	if err = cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (self *productRepository) GetById(ctx context.Context, id string) (*enitity.Product, error) {

	var session = self.mongoClient.NewSession(self.databaseName)

	var result *enitity.Product

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	err = session.Collection(productCollectionName).FindOne(ctx, bson.M{"_id": _id}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (self *productRepository) Insert(ctx context.Context, product *enitity.Product) error {

	var session = self.mongoClient.NewSession(self.databaseName)

	_, err := session.Collection(productCollectionName).InsertOne(ctx, product)

	return err
}

func (self *productRepository) UpdateById(ctx context.Context, id string, product *enitity.Product) error {

	var session = self.mongoClient.NewSession(self.databaseName)

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = session.Collection(productCollectionName).UpdateByID(ctx, _id,
		bson.M{"$set": bson.M{"name": product.Name,
			"categoryId": product.CategoryId,
			"price":      product.Price,
			"updatedAt":  time.Now().UTC(),
		}})

	return err
}

func (self *productRepository) RemoveById(ctx context.Context, id string) error {

	var session = self.mongoClient.NewSession(self.databaseName)

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = session.Collection(productCollectionName).DeleteOne(ctx, bson.M{"_id": _id})

	return err
}

func NewProductRepository(client mongoclient.MongoClient, databaseName string) ProductRepository {
	return &productRepository{mongoClient: client, databaseName: databaseName}
}
