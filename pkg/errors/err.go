package errors

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Message map[string]interface{}

type RespError struct {
	Code    int
	Message Message
}

func (r *RespError) Error() string {
	return fmt.Sprintf("%d: %s", r.Code, r.Message)
}

func defaultMessage(code int, msg string) *RespError {
	return &RespError{
		Code: http.StatusInternalServerError,
		Message: Message{
			"code":  http.StatusInternalServerError,
			"msg":   msg,
			"error": true,
		},
	}
}

func InternalServerError(msgs ...Message) error {
	if len(msgs) == 0 {
		return defaultMessage(http.StatusInternalServerError, "Internal server error!")
	}

	return &RespError{
		Code:    http.StatusInternalServerError,
		Message: msgs[0],
	}
}

func BadRequest(msgs ...Message) error {
	if len(msgs) == 0 {
		return defaultMessage(http.StatusBadRequest, "Ilegal request provided!")
	}

	return &RespError{
		Code:    http.StatusBadRequest,
		Message: msgs[0],
	}
}

func NotFound(msgs ...Message) error {
	if len(msgs) == 0 {
		return defaultMessage(http.StatusNotFound, "Not found!")
	}

	return &RespError{
		Code:    http.StatusNotFound,
		Message: msgs[0],
	}
}

func Unauthorization(msgs ...Message) error {
	if len(msgs) == 0 {
		return defaultMessage(http.StatusUnauthorized, "Unauthorized!")
	}

	return &RespError{
		Code:    http.StatusUnauthorized,
		Message: msgs[0],
	}
}

func ErrorResponse(ctx *fiber.Ctx, err error) error {
	if err, ok := err.(*RespError); ok {
		return ctx.Status(err.Code).JSON(err.Message)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}
