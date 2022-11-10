package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID                 primitive.ObjectID `json:"_id" bson:"_id"`
	Title              *string            `json:"title" validate:"required,min=2,max=100"`
	Description        *string            `json:"description"`
	Price              float32            `json:"price"`
	DiscountPercentage float64            `json:"discountPercentage,truncate"`
	Rating             float64            `json:"rating"`
	Stock              int32              `json:"stock"`
	Brand              *string            `json:"brand"`
	Category           *string            `json:"category"`
	Thumbnail          *string            `json:"thumbnail"`
	Images             []string           `json:"images"`
	CreatedAt          time.Time          `json:"createdAt`
	UpdatedAt          time.Time          `json:"updatedAt"`
}
