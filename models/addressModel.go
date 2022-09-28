package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	UserId    primitive.ObjectID `json:"_user" bson:"_user" `
	Name      *string            `json:"name" validate:"required,min=2,max=100"`
	Address   *string            `json:"address" validate:"required,min=2,max=100"`
	City      *string            `json:"city"`
	State     *string            `json:"state"`
	Email     *string            `json:"email" validate:"email,required"`
	Phone     *string            `json:"phone" validate:"required"`
	CreatedAt time.Time          `json:"createdAt`
	UpdatedAt time.Time          `json:"updatedAt"`
}
