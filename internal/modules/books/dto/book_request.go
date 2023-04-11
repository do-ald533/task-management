package dto

import "github.com/google/uuid"

type BookAttrs struct {
	Picture     string `json:"picture,omitempty"`
	Description string `json:"description,omitempty"`
	Rating      *int   `json:"rating,omitempty" validate:"min=1,max=10"`
}

type BookStoreRequest struct {
	Title     string    `json:"title" validate:"required,lte=255"`
	Author    string    `json:"author" validate:"required,lte=255"`
	BookAttrs BookAttrs `json:"book_attrs" validate:"required,dive"`
}

type BookUpdateRequest struct {
	Title      string     `json:"title,omitempty" validate:"lte=255"`
	Author     string     `json:"author,omitempty" validate:"lte=255"`
	UserID     *uuid.UUID `json:"user_id,omitempty" validate:"omitempty,uuid"`
	BookStatus *int       `json:"book_status,omitempty"`
	BookAttrs  BookAttrs  `json:"book_attrs,omitempty" validate:"dive"`
}

type BookDeleteRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}
