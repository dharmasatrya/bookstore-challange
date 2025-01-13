// service/borrow_service.go
package service

import (
	"borrow/entity"
	"borrow/src/repository"
	"context"
	"fmt"
	"time"

	pb "github.com/dharmasatrya/proto-repo/borrow"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type BorrowService struct {
	pb.UnimplementedBorrowServiceServer
	borrowRepo repository.BorrowRepository
}

// var jwtSecret = []byte("secret")

func NewBorrowService(borrowRepository repository.BorrowRepository) *BorrowService {
	return &BorrowService{
		borrowRepo: borrowRepository,
	}
}

func (s *BorrowService) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {

	_, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	// Get claims from context that was set in auth middleware
	claims, ok := ctx.Value("claims").(jwt.MapClaims)
	if !ok {
		return nil, status.Errorf(codes.Internal, "failed to get user claims")
	}

	// Extract user_id from claims
	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, status.Errorf(codes.Internal, "user_id not found in claims")
	}

	borrowedDate, err := time.Parse("02-01-2006", req.BorrowedDate)
	if err != nil {
		return nil, err
	}

	borrowInput := entity.BorrowedBook{
		BookID:       req.BookId,
		UserID:       userID,
		BorrowedDate: borrowedDate,
		ReturnDate:   nil,
	}

	res, err := s.borrowRepo.BorrowBook(borrowInput)
	if err != nil {
		fmt.Println("error")
	}

	var returnDateStr string
	if res.ReturnDate != nil {
		returnDateStr = res.ReturnDate.Format("02-01-2006")
	}

	return &pb.BorrowBookResponse{
		Id:           res.ID.Hex(),
		BookId:       res.BookID,
		UserId:       res.UserID,
		BorrowedDate: res.BorrowedDate.Format("02-01-2006"),
		ReturnDate:   returnDateStr, // Will be empty string if ReturnDate is nil
	}, nil
}

func (s *BorrowService) EditBorrowedBook(ctx context.Context, req *pb.EditBorrowedBookRequest) (*pb.EditBorrowedBookResponse, error) {

	// Convert ID from request
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	// Create update input
	updateInput := entity.EditBorrowedBookRequest{
		ID: objectID,
	}

	// Handle optional fields
	if req.BookId != "" {
		bookID := req.BookId
		updateInput.BookID = &bookID
	}

	// Parse dates if provided
	if req.BorrowedDate != "" {
		borrowedDate, err := time.Parse("02-01-2006", req.BorrowedDate)
		if err != nil {
			return nil, fmt.Errorf("invalid borrowed date format: %w", err)
		}
		updateInput.BorrowedDate = &borrowedDate
	}

	if req.ReturnDate != "" {
		returnDate, err := time.Parse("02-01-2006", req.ReturnDate)
		if err != nil {
			return nil, fmt.Errorf("invalid return date format: %w", err)
		}
		updateInput.ReturnDate = &returnDate
	}

	// Call repository
	updatedBorrow, err := s.borrowRepo.EditBorrowedBook(ctx, updateInput)
	if err != nil {
		return nil, fmt.Errorf("failed to update borrowed book: %w", err)
	}

	// Format the return date string if it exists
	var returnDateStr string
	if updatedBorrow.ReturnDate != nil {
		returnDateStr = updatedBorrow.ReturnDate.Format("02-01-2006")
	}

	// Transform to response
	return &pb.EditBorrowedBookResponse{
		Id:           updatedBorrow.ID.Hex(),
		BookId:       updatedBorrow.BookID,
		UserId:       updatedBorrow.UserID,
		BorrowedDate: updatedBorrow.BorrowedDate.Format("02-01-2006"),
		ReturnDate:   returnDateStr,
	}, nil
}
