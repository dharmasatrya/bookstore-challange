// service/service_test.go
package service

import (
	"borrow/entity"
	"borrow/src/repository/mock"
	"context"
	"testing"
	"time"

	pb "github.com/dharmasatrya/proto-repo/borrow"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
)

func TestBorrowService_BorrowBook(t *testing.T) {
	// Setup
	mockRepo := new(mock.MockBorrowRepository)
	service := NewBorrowService(mockRepo)

	tests := []struct {
		name          string
		input         *pb.BorrowBookRequest
		setupAuth     func(context.Context) context.Context
		mockBehavior  func()
		expectedError bool
		checkResponse func(*testing.T, *pb.BorrowBookResponse)
	}{
		{
			name: "Success",
			input: &pb.BorrowBookRequest{
				BookId:       "book123",
				BorrowedDate: "13-01-2025",
			},
			setupAuth: func(ctx context.Context) context.Context {
				md := metadata.New(map[string]string{
					"authorization": "Bearer valid-token",
				})
				ctx = metadata.NewIncomingContext(ctx, md)

				claims := jwt.MapClaims{
					"user_id": "user123",
				}
				return context.WithValue(ctx, "claims", claims)
			},
			// service/service_test.go
			mockBehavior: func() {
				borrowedDate, _ := time.Parse("02-01-2006", "13-01-2025")
				expectedInput := entity.BorrowedBook{ // Changed from BorrowBookInput to BorrowedBook
					BookID:       "book123",
					UserID:       "user123",
					BorrowedDate: borrowedDate,
					ReturnDate:   nil,
				}

				mockRepo.On("BorrowBook", expectedInput).Return(&entity.BorrowedBook{
					ID:           primitive.NewObjectID(),
					BookID:       "book123",
					UserID:       "user123",
					BorrowedDate: borrowedDate,
					ReturnDate:   nil,
				}, nil)
			},
			expectedError: false,
			checkResponse: func(t *testing.T, response *pb.BorrowBookResponse) {
				assert.NotEmpty(t, response.Id)
				assert.Equal(t, "book123", response.BookId)
				assert.Equal(t, "user123", response.UserId)
				assert.Equal(t, "13-01-2025", response.BorrowedDate)
				assert.Empty(t, response.ReturnDate)
			},
		},
		{
			name: "Invalid Date Format",
			input: &pb.BorrowBookRequest{
				BookId:       "book123",
				BorrowedDate: "2025-01-13", // wrong format
			},
			setupAuth: func(ctx context.Context) context.Context {
				md := metadata.New(map[string]string{
					"authorization": "Bearer valid-token",
				})
				ctx = metadata.NewIncomingContext(ctx, md)

				claims := jwt.MapClaims{
					"user_id": "user123",
				}
				return context.WithValue(ctx, "claims", claims)
			},
			mockBehavior:  func() {}, // No mock needed as it should fail before repo call
			expectedError: true,
			checkResponse: func(t *testing.T, response *pb.BorrowBookResponse) {
				assert.Nil(t, response)
			},
		},
		{
			name: "Missing Auth",
			input: &pb.BorrowBookRequest{
				BookId:       "book123",
				BorrowedDate: "13-01-2025",
			},
			setupAuth: func(ctx context.Context) context.Context {
				return ctx // no auth added
			},
			mockBehavior:  func() {}, // No mock needed as it should fail before repo call
			expectedError: true,
			checkResponse: func(t *testing.T, response *pb.BorrowBookResponse) {
				assert.Nil(t, response)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup context with auth
			ctx := context.Background()
			ctx = tt.setupAuth(ctx)

			// Setup mock behavior
			tt.mockBehavior()

			// Call service
			response, err := service.BorrowBook(ctx, tt.input)

			// Check error
			if tt.expectedError {
				assert.Error(t, err)
				tt.checkResponse(t, response)
				return
			}

			// Check success case
			assert.NoError(t, err)
			tt.checkResponse(t, response)
		})
	}

	// Verify all expectations were met
	mockRepo.AssertExpectations(t)
}
