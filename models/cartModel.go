package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID       primitive.ObjectID `json:"_id"`
	Quantity int                `json:"quantity"`
	Price    float32            `json:"price"`
}

type Cart struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	UserId    primitive.ObjectID `json:"_user" bson:"_user"`
	Items     []Item             `json:"items"`
	Quantity  int                `josn:"quantity"`
	CreatedAt time.Time          `json:"createdAt`
	UpdatedAt time.Time          `json:"updatedAt"`
}
