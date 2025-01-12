// service/book_service.go
package service

import (
	"books/entity"
	"books/src/repository"
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/dharmasatrya/proto-repo/book"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookService struct {
	pb.UnimplementedBookServiceServer
	bookRepo repository.BookRepository
}

// var jwtSecret = []byte("secret")

func NewBookService(bookRepository repository.BookRepository) *BookService {
	return &BookService{
		bookRepo: bookRepository,
	}
}

func (s *BookService) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	log.Printf("Received : %s", req.Title)

	publishedDate, err := time.Parse("02-01-2006", req.PublishedDate)
	if err != nil {
		return nil, err
	}

	bookInput := entity.Book{
		Title:         req.Title,
		Author:        req.Author,
		PublishedDate: publishedDate,
		Status:        "Available",
		UserId:        "",
	}

	res, err := s.bookRepo.CreateBook(bookInput)
	if err != nil {
		fmt.Println("error")
	}

	return &pb.CreateBookResponse{
		Id:            res.ID.Hex(),
		Title:         res.Title,
		Author:        res.Author,
		PublishedDate: res.PublishedDate.Format("02-01-2006"),
		Status:        res.Status,
		UserId:        res.UserId,
	}, nil
}

// service/book_service.go
func (s *BookService) EditBook(ctx context.Context, req *pb.EditBookRequest) (*pb.EditBookResponse, error) {
	log.Printf("Editing book with ID: %s", req.Id)

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid book ID: %w", err)
	}

	// Parse date if it's provided
	var publishedDate *time.Time
	if req.PublishedDate != nil {
		parsed, err := time.Parse("02-01-2006", *req.PublishedDate)
		if err != nil {
			return nil, fmt.Errorf("invalid date format: %w", err)
		}
		publishedDate = &parsed
	}

	// Create update input
	updateInput := entity.EditBookRequest{
		ID:            objectID,
		Title:         req.Title,
		Author:        req.Author,
		PublishedDate: publishedDate,
		Status:        req.Status,
		UserId:        req.UserId,
	}

	// Call repository
	updatedBook, err := s.bookRepo.EditBook(ctx, updateInput)
	if err != nil {
		return nil, fmt.Errorf("failed to update book: %w", err)
	}

	// Transform to response
	return &pb.EditBookResponse{
		Id:            updatedBook.ID.Hex(),
		Title:         updatedBook.Title,
		Author:        updatedBook.Author,
		PublishedDate: updatedBook.PublishedDate.Format("02-01-2006"),
		Status:        updatedBook.Status,
		UserId:        updatedBook.UserId,
	}, nil
}

func (s *BookService) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	log.Printf("Deleting book with ID: %s", req.Id)

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid book ID: %w", err)
	}

	// Call repository
	deletedBook, err := s.bookRepo.DeleteBook(ctx, objectID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete book: %w", err)
	}

	// Transform to response
	return &pb.DeleteBookResponse{
		Id:            deletedBook.ID.Hex(),
		Title:         deletedBook.Title,
		Author:        deletedBook.Author,
		PublishedDate: deletedBook.PublishedDate.Format("02-01-2006"),
		Status:        deletedBook.Status,
		UserId:        deletedBook.UserId,
	}, nil
}
