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
