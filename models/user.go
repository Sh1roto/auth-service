package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name,omitempty" validate:"required,min=2,max=100"`
	LastName     *string            `json:"last_name,omitempty" validate:"required,min=2,max=100"`
	Password     *string            `json:"password,omitempty" validate:"required,min=6"`
	Email        *string            `json:"email,omitempty" validate:"required,email"`
	Phone        *string            `json:"phone,omitempty" validate:"required"`
	Token        *string            `json:"token,omitempty"`
	UserType     *string            `json:"user_type,omitempty" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken *string            `json:"refresh_token,omitempty"`
	Created_at   time.Time          `json:"created_at"`
	Updated_at   time.Time          `json:"updated_at"`
	UserID       string             `json:"user_id"`
}
