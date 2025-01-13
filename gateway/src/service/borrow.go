package service

import (
	"context"
	"gateway/entity"
	"log"
	"net/http"

	pb "github.com/dharmasatrya/proto-repo/borrow"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type BorrowService interface {
	BorrowBook(token string, input entity.BorrowBookInput) (int, *entity.BorrowedBook)
	EditBorrowedBook(token string, id string, input entity.EditBorrowRequest) (int, *entity.BorrowedBook)
}

type borrowService struct {
	conn *grpc.ClientConn
}

func NewBorrowService(conn *grpc.ClientConn) *borrowService {
	return &borrowService{conn}
}

func (s *borrowService) BorrowBook(token string, input entity.BorrowBookInput) (int, *entity.BorrowedBook) {
	client := pb.NewBorrowServiceClient(s.conn)

	md := metadata.Pairs("authorization", token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.BorrowBook(ctx, &pb.BorrowBookRequest{
		BookId:       input.BookID,
		BorrowedDate: input.BorrowedDate,
	})

	if err != nil {
		log.Printf("error while borrowing book: %v", err)
		return http.StatusInternalServerError, nil
	}

	response := &entity.BorrowedBook{
		ID:           res.Id,
		BookID:       res.BookId,
		UserID:       res.UserId,
		BorrowedDate: res.BorrowedDate,
		ReturnDate:   res.ReturnDate,
	}

	return http.StatusOK, response
}

func (s *borrowService) EditBorrowedBook(token string, id string, input entity.EditBorrowRequest) (int, *entity.BorrowedBook) {
	client := pb.NewBorrowServiceClient(s.conn)

	md := metadata.Pairs("authorization", token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	req := &pb.EditBorrowedBookRequest{
		Id: id,
	}

	if input.BookID != nil {
		req.BookId = *input.BookID
	}
	if input.BorrowedDate != nil {
		req.BorrowedDate = *input.BorrowedDate
	}
	if input.ReturnDate != nil {
		req.ReturnDate = *input.ReturnDate
	}

	res, err := client.EditBorrowedBook(ctx, req)
	if err != nil {
		log.Printf("error while editing borrowed book: %v", err)
		return http.StatusInternalServerError, nil
	}

	response := &entity.BorrowedBook{
		ID:           res.Id,
		BookID:       res.BookId,
		UserID:       res.UserId,
		BorrowedDate: res.BorrowedDate,
		ReturnDate:   res.ReturnDate,
	}

	return http.StatusOK, response
}
