package mocks

import (
	"goapi/internal/modules/books/dto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type BookServiceMock struct {
	mock.Mock
}

func (m *BookServiceMock) Store(userId uuid.UUID, body *dto.BookStoreRequest) (*dto.BookResponse, error) {
	args := m.Called(userId, body)
	return args.Get(0).(*dto.BookResponse), args.Error(1)
}

func (m *BookServiceMock) Update(id *uuid.UUID, body *dto.BookUpdateRequest) (*dto.BookResponse, error) {
	args := m.Called(id, body)
	return args.Get(0).(*dto.BookResponse), args.Error(1)
}

func (m *BookServiceMock) Show(id uuid.UUID) (*dto.BookResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.BookResponse), args.Error(1)
}

func (m *BookServiceMock) Index(limit, page int64) (*dto.BookIndexResponse, error) {
	args := m.Called(limit, page)
	return args.Get(0).(*dto.BookIndexResponse), args.Error(1)
}

func (m *BookServiceMock) Delete(id *uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
