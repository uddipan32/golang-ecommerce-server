package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson: "_id,omitempty"`
	Name        *string            `json:"name" validate:"required,min=2,max=100"`
	Description *string            `json:"description"`
	price       float32            `json:"price"`
	Images      []string           `json:"images"`
	CreatedAt   time.Time          `json:"createdAt`
	UpdatedAt   time.Time          `json:"updatedAt"`
}
