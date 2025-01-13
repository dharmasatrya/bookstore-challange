// repository/mock/borrow_repository_mock.go
package mock

import (
	"borrow/entity"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockBorrowRepository struct {
	mock.Mock
}

func (m *MockBorrowRepository) BorrowBook(input entity.BorrowedBook) (*entity.BorrowedBook, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.BorrowedBook), args.Error(1)
}

func (m *MockBorrowRepository) EditBorrowedBook(ctx context.Context, input entity.EditBorrowedBookRequest) (*entity.BorrowedBook, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.BorrowedBook), args.Error(1)
}
