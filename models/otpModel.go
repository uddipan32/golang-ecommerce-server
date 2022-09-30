package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OTP struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Phone     *string            `json:"phone" validate:"required"`
	OTP       *string            `json:"otp"`
	CreatedAt time.Time          `json:"createdAt`
	UpdatedAt time.Time          `json:"updatedAt"`
}
