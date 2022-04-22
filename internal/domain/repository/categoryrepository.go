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

const categoryCollectionName = "category"

type CategoryRepository interface {
	GetAll(ctx context.Context) ([]enitity.Category, error)
	GetById(ctx context.Context, id string) (*enitity.Category, error)
	Insert(ctx context.Context, category *enitity.Category) error
	UpdateById(ctx context.Context, id string, category *enitity.Category) error
	RemoveById(ctx context.Context, id string) error
}

type categoryRepository struct {
	mongoClient  mongoclient.MongoClient
	databaseName string
}

func (self *categoryRepository) GetAll(ctx context.Context) ([]enitity.Category, error) {

	var session = self.mongoClient.NewSession(self.databaseName)

	var result []enitity.Category

	cur, err := session.Collection(categoryCollectionName).Find(ctx, bson.D{})

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

func (self *categoryRepository) GetById(ctx context.Context, id string) (*enitity.Category, error) {

	var session = self.mongoClient.NewSession(self.databaseName)

	var result *enitity.Category

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	err = session.Collection(categoryCollectionName).FindOne(ctx, bson.M{"_id": _id}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (self *categoryRepository) Insert(ctx context.Context, category *enitity.Category) error {

	var session = self.mongoClient.NewSession(self.databaseName)

	_, err := session.Collection(categoryCollectionName).InsertOne(ctx, category)

	return err
}

func (self *categoryRepository) UpdateById(ctx context.Context, id string, category *enitity.Category) error {

	var session = self.mongoClient.NewSession(self.databaseName)

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = session.Collection(categoryCollectionName).UpdateByID(ctx, _id, bson.M{"$set": bson.M{"name": category.Name, "updatedAt": time.Now().UTC()}})

	return err
}

func (self *categoryRepository) RemoveById(ctx context.Context, id string) error {

	var session = self.mongoClient.NewSession(self.databaseName)

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = session.Collection(categoryCollectionName).DeleteOne(ctx, bson.M{"_id": _id})

	return err
}

func NewCategoryRepository(client mongoclient.MongoClient, databaseName string) CategoryRepository {
	return &categoryRepository{mongoClient: client, databaseName: databaseName}
}
