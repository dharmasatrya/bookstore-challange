package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"name" bson:"name"`
	Password string             `json:"email" bson:"email"`
}

type CreateUserInput struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}
