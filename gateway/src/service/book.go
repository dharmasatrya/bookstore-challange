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
