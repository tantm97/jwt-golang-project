package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name" binding:"required,min=2,max=100"`
	LastName     *string            `json:"last_name" binding:"required,min=2,max=100"`
	Password     *string            `json:"password" binding:"required,min=6"`
	Email        *string            `json:"email" binding:"email,required"`
	Phone        *string            `json:"phone" binding:"required"`
	Token        *string            `json:"token"`
	UserType     *string            `json:"user_type" binding"required,eq=ADMIN|eq=USER"`
	RefreshToken *string            `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdateAt     time.Time          `json:"updated_at"`
	UserId       string             `json:"user_id"`
}
