package repository

import (
	"books/entity"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository interface {
	CreateBook(book entity.Book) (*entity.Book, error)
}

type bookRepository struct {
	db *mongo.Collection
}

func NewBookRepository(db *mongo.Collection) *bookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) CreateBook(input entity.Book) (*entity.Book, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	book := entity.Book{
		ID:            primitive.NewObjectID(),
		Title:         input.Title,
		Author:        input.Author,
		PublishedDate: input.PublishedDate,
		Status:        input.Status,
		UserId:        input.UserId,
	}

	_, err2 := r.db.InsertOne(ctx, book)
	if err2 != nil {
		return nil, err2
	}

	return &book, nil
}
