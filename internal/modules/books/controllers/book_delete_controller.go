package controllers

import (
	"goapi/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Description Delete book by given ID.
// @Summary delete book by given ID
// @Tags Books
// @Accept json
// @Produce json
// @Param request body dto.BookDeleteRequest true "Book ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /api/v1/books [delete]
func (ctrl *BooksController) Delete(req *fiber.Ctx) error {
	id, err := uuid.Parse(req.Params("id"))
	if err != nil {
		return req.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := ctrl.bookServices.Delete(&id); err != nil {
		return errors.ErrorResponse(req, err)
	}

	return req.Status(fiber.StatusNoContent).JSON(nil)
}
