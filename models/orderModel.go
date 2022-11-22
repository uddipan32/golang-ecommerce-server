package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Items struct {
	ID       primitive.ObjectID `json:"_id"`
	Quantity int                `json:"quantity"`
	Price    float32            `json:"price"`
}

type Order struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	Items         []Item             `json:"items"`
	Quantity      int                `josn:"quantity"`
	TotalPrice    float32            `json:"totalPrice"`
	Status        *string            `json:"status"`
	PaymentStatus *string            `json:"paymentStatus"`
	PaymentMode   *string            `json:"paymentMode"`
	CreatedAt     time.Time          `json:"createdAt`
	UpdatedAt     time.Time          `json:"updatedAt"`
}
