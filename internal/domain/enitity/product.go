package enitity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	CategoryId string             `json:"categoryId" bson:"categoryId"`
	Name       string             `json:"name" bson:"name"`
	Price      float64            `json:"price" bson:"price"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func NewProduct(categoryId, name string, price float64) *Product {
	return &Product{
		Id:         primitive.NewObjectID(),
		CategoryId: categoryId,
		Name:       name,
		Price:      price,
		CreatedAt:  time.Now().UTC(),
	}
}
