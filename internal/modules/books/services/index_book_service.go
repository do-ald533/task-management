package services

import (
	"goapi/internal/modules/books/dto"
	"goapi/internal/modules/books/entities"
	"goapi/pkg/pagination"
)

func (serv *BookService) Index(limit, page int64) (*dto.BookIndexResponse, error) {
	// Get all allBooks.
	books, count, err := serv.BookRepository.Index(limit, page)

	result := pagination.Paginate[entities.Book]{
		Items:         books,
		MongoPaginate: pagination.NewMongoPaginate(limit, page, count),
	}

	if err != nil {
		// Return, if books not found.
		return nil, err
	}

	// Return status 200 OK.
	return &dto.BookIndexResponse{Paginate: &result}, nil
}
