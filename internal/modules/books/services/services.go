package services

import (
	"goapi/internal/modules/books/dto"
	"goapi/internal/modules/books/repositories"
	"goapi/pkg/aws/sns"

	"github.com/google/uuid"
)

type BookServicesImpl interface {
	Store(userId uuid.UUID, body *dto.BookStoreRequest) (*dto.BookResponse, error)
	Show(id uuid.UUID) (*dto.BookResponse, error)
	Index(limit, page int64) (*dto.BookIndexResponse, error)
	Update(id *uuid.UUID, body *dto.BookUpdateRequest) (*dto.BookResponse, error)
	Delete(id *uuid.UUID) error
}

type BookService struct {
	BookRepository repositories.BookRepositoryImpl
	SnsService     sns.SNSService
}

func NewBookService(
	bookRepository repositories.BookRepositoryImpl,
	snsService sns.SNSService,
) BookServicesImpl {

	return &BookService{bookRepository, snsService}
}
