package service

import (
	"context"
	"gateway/entity"
	"log"
	"net/http"

	bookConn "github.com/dharmasatrya/proto-repo/book"
	pb "github.com/dharmasatrya/proto-repo/book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type BookService interface {
	CreateBook(token string, input entity.CreateBookInput) (int, *entity.Book)
	EditBook(token string, id string, input entity.EditBookRequest) (int, *entity.Book)
	DeleteBook(token string, id string) (int, *entity.Book)
}

type bookService struct {
	conn *grpc.ClientConn
}

func NewBookService(conn *grpc.ClientConn) *bookService {
	return &bookService{conn}
}

func (u *bookService) CreateBook(token string, input entity.CreateBookInput) (int, *entity.Book) {
	client := bookConn.NewBookServiceClient(u.conn)

	md := metadata.Pairs("authorization", token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.CreateBook(ctx, &pb.CreateBookRequest{
		Title:         input.Title,
		Author:        input.Author,
		PublishedDate: input.PublishedDate,
		Status:        "",
		UserId:        "",
	})

	if err != nil {
		log.Fatalf("error while create request %v", err)
	}

	response := entity.Book{
		ID:            res.Id,
		Title:         res.Title,
		Author:        res.Author,
		PublishedDate: res.PublishedDate,
		Status:        res.Status,
		UserId:        res.UserId,
	}

	return http.StatusOK, &response
}

func (u *bookService) EditBook(token string, id string, input entity.EditBookRequest) (int, *entity.Book) {
	client := bookConn.NewBookServiceClient(u.conn)

	md := metadata.Pairs("authorization", token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	req := &pb.EditBookRequest{
		Id: id,
	}

	if input.Title != nil {
		req.Title = input.Title
	}
	if input.Author != nil {
		req.Author = input.Author
	}
	if input.PublishedDate != nil {
		req.PublishedDate = input.PublishedDate
	}
	if input.Status != nil {
		req.Status = input.Status
	}
	if input.UserId != nil {
		req.UserId = input.UserId
	}

	res, err := client.EditBook(ctx, req)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	response := &entity.Book{
		ID:            res.Id,
		Title:         res.Title,
		Author:        res.Author,
		PublishedDate: res.PublishedDate,
		Status:        res.Status,
		UserId:        res.UserId,
	}

	return http.StatusOK, response
}

func (u *bookService) DeleteBook(token string, id string) (int, *entity.Book) {
	client := bookConn.NewBookServiceClient(u.conn)

	md := metadata.Pairs("authorization", token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	req := &pb.DeleteBookRequest{
		Id: id,
	}

	res, err := client.DeleteBook(ctx, req)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	response := &entity.Book{
		ID:            res.Id,
		Title:         res.Title,
		Author:        res.Author,
		PublishedDate: res.PublishedDate,
		Status:        res.Status,
		UserId:        res.UserId,
	}

	return http.StatusOK, response
}
