package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      *string            `json:"name" validate:"required,min=2,max=100"`
	Email     *string            `json:"email" validate:"email,required"`
	Phone     *string            `json:"phone" validate:"required"`
	OTP       *string            `json:"otp"`
	CreatedAt time.Time          `json:"createdAt`
	UpdatedAt time.Time          `json:"updatedAt"`
}
