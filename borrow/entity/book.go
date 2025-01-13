package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BorrowedBook struct {
	ID           primitive.ObjectID `json:"id" bson:"id"`
	BookID       string             `json:"book_id" bson:"book_id"`
	UserID       string             `json:"user_id" bson:"user_id"`
	BorrowedDate time.Time          `json:"borrowed_date" bson:"borrowed_date"`
	ReturnDate   *time.Time         `json:"return_date,omitempty" bson:"return_date"`
}

type BorrowBookInput struct {
	ID           primitive.ObjectID `json:"id" bson:"id"`
	BookID       string             `json:"book_id" bson:"book_id"`
	UserID       string             `json:"user_id" bson:"user_id"`
	BorrowedDate string             `json:"borrowed_date" bson:"borrowed_date" validate:"required,datetime=2006-01-02"`
}

type EditBorrowedBookRequest struct {
	ID           primitive.ObjectID `json:"id" bson:"id"`
	BookID       *string            `json:"book_id,omitempty" bson:"book_id"`
	UserID       *string            `json:"user_id,omitempty" bson:"user_id"`
	BorrowedDate *time.Time         `json:"borrowed_date,omitempty" bson:"borrowed_date"`
	ReturnDate   *time.Time         `json:"return_date,omitempty" bson:"return_date"`
}
