package controllers

import (
	"goapi/internal/modules/books/services"
)

type BooksController struct {
	bookServices services.BookServicesImpl
}

func NewBooksController(
	bookServices services.BookServicesImpl,
) *BooksController {
	return &BooksController{bookServices}
}
