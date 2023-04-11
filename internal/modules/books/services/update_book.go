package services

import (
	"goapi/internal/modules/books/dto"
	"goapi/internal/modules/books/entities"
	"goapi/pkg/convert"
	"time"

	"github.com/google/uuid"
)

func (repo *BookService) Update(id *uuid.UUID, body *dto.BookUpdateRequest) (*dto.BookResponse, error) {
	book := entities.Book{}
	// Parse from dto to entities struct
	convert.ToStruct(*body, &book)

	now := time.Now()
	book.UpdatedAt = &now

	// Get all books.
	books, err := repo.BookRepository.Update(id, &book)
	if err != nil {
		// Return, if books not found.
		return nil, err
	}

	res := dto.BookResponse{}
	convert.ToStruct(*books, &res)

	// Return status 200 OK.
	return &res, nil
}
