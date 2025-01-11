package entity

import (
	"time"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	Success      bool
	ErrorMessage string
}

type Book struct {
	ID            string `json:"id" bson:"id"`
	Title         string `json:"title" bson:"title"`
	Author        string `json:"author" bson:"author"`
	PublishedDate string `json:"published_date" bson:"published_date"`
	Status        string `json:"status" bson:"status"`
	UserId        string `json:"user_id" bson:"user_id"`
}

type CreateBookInput struct {
	Title         string `json:"title" bson:"title"`
	Author        string `json:"author" bson:"author"`
	PublishedDate string `json:"published_date" bson:"published_date" validate:"required,datetime=2006-01-02"`
}

type CreateBookResponse struct {
	ID            string    `json:"id" bson:"id"`
	Title         string    `json:"title" bson:"title"`
	Author        string    `json:"author" bson:"author"`
	PublishedDate time.Time `json:"published_date" bson:"published_date"`
	Status        string    `json:"status" bson:"status"`
	UserId        string    `json:"user_id" bson:"user_id"`
}

type EditBookRequest struct {
	Title         *string `json:"title,omitempty" bson:"title"`
	Author        *string `json:"author,omitempty" bson:"author"`
	PublishedDate *string `json:"published_date,omitempty" bson:"published_date" validate:"required,datetime=2006-01-02"`
	Status        *string `json:"status,omitempty" bson:"status"`
	UserId        *string `json:"user_id,omitempty" bson:"user_id"`
}

type EditBookResponse struct {
	ID            string    `json:"id" bson:"id"`
	Title         string    `json:"title" bson:"title"`
	Author        string    `json:"author" bson:"author"`
	PublishedDate time.Time `json:"published_date" bson:"published_date"`
	Status        string    `json:"status" bson:"status"`
	UserId        string    `json:"user_id" bson:"user_id"`
}

type DeleteBookResponse struct {
	ID            string    `json:"id" bson:"id"`
	Title         string    `json:"title" bson:"title"`
	Author        string    `json:"author" bson:"author"`
	PublishedDate time.Time `json:"published_date" bson:"published_date"`
	Status        string    `json:"status" bson:"status"`
	UserId        string    `json:"user_id" bson:"user_id"`
}
