package enitity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Category struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func NewCategory(name string) *Category {
	return &Category{
		Id:        primitive.NewObjectID(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}
}
