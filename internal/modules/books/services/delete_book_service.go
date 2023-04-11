package services

import (
	"github.com/google/uuid"
)

func (serv *BookService) Delete(id *uuid.UUID) error {
	err := serv.BookRepository.Delete(*id)

	if err != nil {
		return err
	}

	return nil
}
