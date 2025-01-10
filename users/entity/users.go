package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type CreateUserInput struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}
