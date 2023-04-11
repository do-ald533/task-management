package services

import (
	"goapi/internal/modules/books/dto"
	"goapi/internal/modules/books/entities"
	"goapi/pkg/aws/sns/snsactions"
	"goapi/pkg/convert"
	"os"
	"time"

	"github.com/google/uuid"
)

func (serv BookService) Store(userId uuid.UUID, body *dto.BookStoreRequest) (*dto.BookResponse, error) {
	book := entities.Book{}
	convert.ToStruct(body, &book)

	id := uuid.New()
	status := 1
	now := time.Now()

	book.ID = &id
	book.CreatedAt = &now
	book.BookStatus = &status
	book.UserID = &userId

	newBook, err := serv.BookRepository.Store(&book)
	if err != nil {
		return nil, err

	}

	response := dto.BookResponse{}
	convert.ToStruct(newBook, &response)

	serv.SnsService.Setup(os.Getenv("SNS_CREATED_BOOKS"))
	serv.SnsService.Publish(&response, snsactions.CREATION)

	return &response, nil
}
