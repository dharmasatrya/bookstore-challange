package repository

import (
	"borrow/entity"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BorrowRepository interface {
	BorrowBook(input entity.BorrowedBook) (*entity.BorrowedBook, error)
	EditBorrowedBook(ctx context.Context, input entity.EditBorrowedBookRequest) (*entity.BorrowedBook, error)
}

type borrowRepository struct {
	db *mongo.Collection
}

func NewBorrowRepository(db *mongo.Collection) *borrowRepository {
	return &borrowRepository{db}
}

func (r *borrowRepository) BorrowBook(input entity.BorrowedBook) (*entity.BorrowedBook, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	borrow := entity.BorrowedBook{
		ID:           primitive.NewObjectID(),
		BookID:       input.BookID,
		UserID:       input.UserID,
		BorrowedDate: time.Now(),
		ReturnDate:   nil,
	}

	_, err2 := r.db.InsertOne(ctx, borrow)
	if err2 != nil {
		return nil, err2
	}

	return &borrow, nil
}
func (r *borrowRepository) EditBorrowedBook(ctx context.Context, input entity.EditBorrowedBookRequest) (*entity.BorrowedBook, error) {
	// Create update document
	update := bson.M{}

	if input.BookID != nil {
		update["book_id"] = *input.BookID
	}
	if input.UserID != nil {
		update["user_id"] = *input.UserID
	}
	if input.BorrowedDate != nil {
		update["borrowed_date"] = *input.BorrowedDate
	}
	if input.ReturnDate != nil {
		update["return_date"] = *input.ReturnDate
	}

	// If no fields to update
	if len(update) == 0 {
		// Fetch and return current document
		var borrow entity.BorrowedBook
		err := r.db.FindOne(ctx, bson.M{"id": input.ID}).Decode(&borrow)
		return &borrow, err
	}

	// Add updated_at timestamp
	update["updated_at"] = time.Now()

	// Create the update document
	updateDoc := bson.M{"$set": update}

	// Find and update
	var updatedBorrow entity.BorrowedBook
	err := r.db.FindOneAndUpdate(
		ctx,
		bson.M{"id": input.ID},
		updateDoc,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedBorrow)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("borrowed book not found")
		}
		return nil, fmt.Errorf("failed to update borrowed book: %w", err)
	}

	return &updatedBorrow, nil
}
