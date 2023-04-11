package tokens

import (
	"goapi/internal/modules/tokens/dto"
	"goapi/internal/modules/tokens/services"
	"goapi/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

// CreateToken func create a new token based on the secret provided.
// @Description Create a new a new token based on the secret provided. If provided a exist secret you will create a token with selected role.
// @Summary Create a new token.
// @Tags Token
// @Accept json
// @Produce json
// @Param request body dto.TokenRequest true "Request Body"
// @Success 201 {object} dto.TokenResponse
// @Router /api/v1/tokens [post]
func CreateToken(c *fiber.Ctx) error {
	service, err := services.NewTokenService()
	if err != nil {
		return errors.ErrorResponse(c, err)
	}

	// body := &dto.TokenRequest
	body := &dto.TokenRequest{}

	if err := c.BodyParser(body); err != nil {
		return errors.ErrorResponse(c, err)
	}

	token, err := service.CreateToken(body)
	if err != nil {
		return errors.ErrorResponse(c, err)
	}

	// Return status 201 Created.
	return c.Status(fiber.StatusCreated).JSON(*token)
}
