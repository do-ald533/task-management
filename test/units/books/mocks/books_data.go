package mocks

import (
	"goapi/internal/modules/books/dto"
	"goapi/pkg/jwt"

	"github.com/google/uuid"
)

var (
	idMock = uuid.New()
	rating = 8
)

var UserId = uuid.New()
var UserLogged, _ = jwt.GenerateNewTokens(UserId.String(), "admin", []string{"book:create", "book:update", "book:delete"})

var StoreBookMock = dto.BookStoreRequest{
	Title:  "Title",
	Author: "Jaum",
	BookAttrs: dto.BookAttrs{
		Picture:     "http://image.png",
		Description: "Description",
		Rating:      &rating,
	},
}

var BookMock = dto.BookResponse{
	ID:     &idMock,
	Title:  "Title",
	Author: "Jaum",
	UserID: &UserId,
	BookAttrs: &dto.BookAttrs{
		Picture:     "http://image.png",
		Description: "Description",
		Rating:      &rating,
	},
}
