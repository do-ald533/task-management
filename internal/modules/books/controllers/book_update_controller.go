package controllers

import (
	"goapi/internal/modules/books/dto"
	"goapi/pkg/errors"
	"goapi/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Description Update book.
// @Summary update book
// @Tags Books
// @Accept json
// @Produce json
// @Param request body dto.BookUpdateRequest true "Request Body"
// @Success 202 {string} status "ok"
// @Security ApiKeyAuth
// @Router /api/v1/books [put]
func (ctrl *BooksController) Update(req *fiber.Ctx) error {
	id, err := uuid.Parse(req.Params("id"))
	if err != nil {
		return req.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	body := &dto.BookUpdateRequest{}
	req.BodyParser(body)

	validate := utils.NewValidator()

	if err := validate.Struct(body); err != nil {
		return req.Status(fiber.StatusUnprocessableEntity).JSON(utils.ValidatorErrors(err))
	}

	res, err := ctrl.bookServices.Update(&id, body)
	if err != nil {
		return errors.ErrorResponse(req, err)
	}

	return req.Status(fiber.StatusCreated).JSON(res)
}
