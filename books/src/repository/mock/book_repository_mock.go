// repository/mock/book_repository_mock.go
package mock

import (
	"books/entity"
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) CreateBook(book entity.Book) (*entity.Book, error) {
	args := m.Called(book)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Book), args.Error(1)
}

// Need to implement other methods of the interface
func (m *MockBookRepository) EditBook(ctx context.Context, input entity.EditBookRequest) (*entity.Book, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Book), args.Error(1)
}

func (m *MockBookRepository) DeleteBook(ctx context.Context, id primitive.ObjectID) (*entity.Book, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Book), args.Error(1)
}

func (m *MockBookRepository) GetAllBooks(ctx context.Context) ([]entity.Book, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Book), args.Error(1)
}

func (m *MockBookRepository) GetBookById(ctx context.Context, id primitive.ObjectID) (*entity.Book, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Book), args.Error(1)
}
