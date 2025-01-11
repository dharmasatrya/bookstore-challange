package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
}

type CreateUserInput struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token        string
	Success      bool
	ErrorMessage string
}
