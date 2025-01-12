package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID            primitive.ObjectID `json:"id" bson:"id"`
	Title         string             `json:"title" bson:"title"`
	Author        string             `json:"author" bson:"author"`
	PublishedDate time.Time          `json:"published_date" bson:"published_date"`
	Status        string             `json:"status" bson:"status"`
	UserId        string             `json:"user_id" bson:"user_id"`
}

type CreateBookInput struct {
	Title         string `json:"title" bson:"title"`
	Author        string `json:"author" bson:"author"`
	PublishedDate string `json:"published_date" bson:"published_date" validate:"required,datetime=2006-01-02"`
}

type CreateBookResponse struct {
	ID            primitive.ObjectID `json:"id" bson:"id"`
	Title         string             `json:"title" bson:"title"`
	Author        string             `json:"author" bson:"author"`
	PublishedDate time.Time          `json:"published_date" bson:"published_date"`
	Status        string             `json:"status" bson:"status"`
	UserId        string             `json:"user_id" bson:"user_id"`
}

type EditBookRequest struct {
	ID            primitive.ObjectID `json:"id" bson:"id"`
	Title         *string            `json:"title,omitempty" bson:"title"`
	Author        *string            `json:"author,omitempty" bson:"author"`
	PublishedDate *time.Time         `json:"published_date,omitempty" bson:"published_date"`
	Status        *string            `json:"status,omitempty" bson:"status"`
	UserId        *string            `json:"user_id,omitempty" bson:"user_id"`
}

type EditBookResponse struct {
	ID            primitive.ObjectID `json:"id" bson:"id"`
	Title         string             `json:"title" bson:"title"`
	Author        string             `json:"author" bson:"author"`
	PublishedDate time.Time          `json:"published_date" bson:"published_date"`
	Status        string             `json:"status" bson:"status"`
	UserId        string             `json:"user_id" bson:"user_id"`
}

type DeleteBookResponse struct {
	ID            primitive.ObjectID `json:"id" bson:"id"`
	Title         string             `json:"title" bson:"title"`
	Author        string             `json:"author" bson:"author"`
	PublishedDate time.Time          `json:"published_date" bson:"published_date"`
	Status        string             `json:"status" bson:"status"`
	UserId        string             `json:"user_id" bson:"user_id"`
}
