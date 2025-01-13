// service/service_test.go
package service

import (
	"books/entity"
	"books/src/repository/mock"
	"context"
	"testing"
	"time"

	pb "github.com/dharmasatrya/proto-repo/book"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
)

func TestBookService_CreateBook(t *testing.T) {
	// Setup mock repository
	mockRepo := new(mock.MockBookRepository)
	service := NewBookService(mockRepo)

	// Test data
	publishedDate, _ := time.Parse("02-01-2006", "13-01-2025")
	objectID := primitive.NewObjectID()

	// Setup test case
	t.Run("Success Create Book", func(t *testing.T) {
		// Setup request
		req := &pb.CreateBookRequest{
			Title:         "The Go Programming Language",
			Author:        "Alan A. A. Donovan",
			PublishedDate: "13-01-2025",
			Status:        "Available",
		}

		// Setup auth context
		md := metadata.New(map[string]string{
			"authorization": "Bearer valid-token",
		})
		ctx := metadata.NewIncomingContext(context.Background(), md)

		// Add claims to context
		claims := jwt.MapClaims{
			"user_id": "user123",
		}
		ctx = context.WithValue(ctx, "claims", claims)

		// Setup mock expectations
		expectedInput := entity.Book{
			Title:         req.Title,
			Author:        req.Author,
			PublishedDate: publishedDate,
			Status:        "Available",
			UserId:        "",
		}

		mockRepo.On("CreateBook", expectedInput).Return(&entity.Book{
			ID:            objectID,
			Title:         req.Title,
			Author:        req.Author,
			PublishedDate: publishedDate,
			Status:        "Available",
			UserId:        "",
		}, nil)

		// Call service
		response, err := service.CreateBook(ctx, req)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, objectID.Hex(), response.Id)
		assert.Equal(t, req.Title, response.Title)
		assert.Equal(t, req.Author, response.Author)
		assert.Equal(t, "13-01-2025", response.PublishedDate)
		assert.Equal(t, "Available", response.Status)
		assert.Empty(t, response.UserId)

		// Verify mock
		mockRepo.AssertExpectations(t)
	})
}
