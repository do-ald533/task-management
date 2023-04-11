package controllers

import (
	"goapi/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Description Get book by given ID.
// @Summary get book by given ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} dto.BookResponse
// @Security ApiKeyAuth
// @Router /api/v1/books/{id} [get]
func (ctrl *BooksController) Show(req *fiber.Ctx) error {
	id, err := uuid.Parse(req.Params("id"))
	if err != nil {
		return req.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	book, err := ctrl.bookServices.Show(id)
	if err != nil {
		return errors.ErrorResponse(req, err)
	}

	return req.Status(fiber.StatusOK).JSON(book)
}
