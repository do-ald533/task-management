package controllers

import (
	"goapi/pkg/errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Param limit query int false "limit" minimum(1)
// @Param page query int false "page" minimum(1)
// @Success 200 {object} dto.BookIndexResponse
// @Security ApiKeyAuth
// @Router /api/v1/books [get]
func (ctrl *BooksController) Index(req *fiber.Ctx) error {
	limit, _ := strconv.ParseInt(req.Query("limit", "10"), 10, 64)
	page, _ := strconv.ParseInt(req.Query("page", "1"), 10, 64)

	books, err := ctrl.bookServices.Index(limit, page)
	if err != nil {
		return errors.ErrorResponse(req, err)
	}

	return req.Status(fiber.StatusOK).JSON(*books)
}
