package mocks

import (
	"goapi/internal/modules/books/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	mock.Mock
}

func (m *BookRepositoryMock) Index(limit, page int64) ([]entities.Book, int64, error) {
	args := m.Called(limit, page)
	return args.Get(0).([]entities.Book), args.Get(1).(int64), args.Error(2)
}

func (m *BookRepositoryMock) Show(id uuid.UUID) (*entities.Book, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Book), args.Error(1)
}

func (m *BookRepositoryMock) Store(b *entities.Book) (*entities.Book, error) {
	args := m.Called(b)
	return args.Get(0).(*entities.Book), args.Error(1)
}

func (m *BookRepositoryMock) Update(id *uuid.UUID, book *entities.Book) (*entities.Book, error) {
	args := m.Called(id, book)
	return args.Get(0).(*entities.Book), args.Error(1)
}

func (m *BookRepositoryMock) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
