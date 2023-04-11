package services

import (
	"goapi/internal/modules/books/dto"
	"goapi/pkg/convert"

	"github.com/google/uuid"
)

func (serv *BookService) Show(id uuid.UUID) (*dto.BookResponse, error) {
	// Get all books.
	books, err := serv.BookRepository.Show(id)
	if err != nil {
		// Return, if books not found.
		return nil, err
	}

	res := dto.BookResponse{}

	err = convert.ToStruct(*books, &res)
	if err != nil {
		// Return, if books not found.
		return nil, err
	}

	// Return status 200 OK.
	return &res, nil

}
