package controllers

import (
	"goapi/internal/modules/books/dto"
	"goapi/pkg/errors"
	"goapi/pkg/jwt"
	"goapi/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// @Description Create a new book.
// @Summary create a new book
// @Tags Books
// @Accept json
// @Produce json
// @Param request body dto.BookStoreRequest true "Request Body"
// @Success 200 {object} dto.BookResponse
// @Security ApiKeyAuth
// @Router /api/v1/books [post]
func (ctrl *BooksController) Store(req *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(req)
	if err != nil {
		return errors.ErrorResponse(req, err)
	}

	book := dto.BookStoreRequest{}
	req.BodyParser(&book)

	validate := utils.NewValidator()

	if err := validate.Struct(book); err != nil {
		return req.Status(fiber.StatusUnprocessableEntity).JSON(utils.ValidatorErrors(err))
	}

	response, err := ctrl.bookServices.Store(claims.UserID, &book)
	if err != nil {
		return errors.ErrorResponse(req, err)
	}

	return req.Status(fiber.StatusCreated).JSON(response)
}
