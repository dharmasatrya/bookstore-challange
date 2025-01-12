package repository

import (
	"books/entity"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BookRepository interface {
	CreateBook(book entity.Book) (*entity.Book, error)
	EditBook(ctx context.Context, input entity.EditBookRequest) (*entity.Book, error)
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

// Implement the function
func (r *bookRepository) EditBook(ctx context.Context, input entity.EditBookRequest) (*entity.Book, error) {
	// Create update document
	update := bson.M{}

	if input.Title != nil {
		update["title"] = *input.Title
	}
	if input.Author != nil {
		update["author"] = *input.Author
	}
	if input.PublishedDate != nil {
		update["published_date"] = *input.PublishedDate
	}
	if input.Status != nil {
		update["status"] = *input.Status
	}
	if input.UserId != nil {
		update["user_id"] = *input.UserId
	}

	// If no fields to update
	if len(update) == 0 {
		// Fetch and return current document
		var book entity.Book
		err := r.db.FindOne(ctx, bson.M{"id": input.ID}).Decode(&book)
		return &book, err
	}

	// Add updated_at timestamp
	update["updated_at"] = time.Now()

	// Create the update document
	updateDoc := bson.M{"$set": update}

	// Find and update
	var updatedBook entity.Book
	err := r.db.FindOneAndUpdate(
		ctx,
		bson.M{"id": input.ID},
		updateDoc,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedBook)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("book not found")
		}
		return nil, fmt.Errorf("failed to update book: %w", err)
	}

	return &updatedBook, nil
}
