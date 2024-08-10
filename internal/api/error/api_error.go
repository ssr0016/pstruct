package error

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"task-management-system/internal/api/model"
	"task-management-system/internal/api/response"
)

func NewApiError(err error, code int, message string, data interface{}) *model.ApiError {
	log.Printf("error: %v\n", err)
	return &model.ApiError{
		Code:    code,
		Message: message,
		Data:    &data,
	}
}

func DefaultErrorHandler(ctx *fiber.Ctx, err error) error {
	e, ok := err.(*model.ApiError)
	if !ok {
		ef, ok := err.(*fiber.Error)
		if !ok {
			e = NewApiError(err, fiber.StatusInternalServerError, err.Error(), nil)
		} else {
			e = NewApiError(err, ef.Code, ef.Error(), nil)
		}
	}

	return ctx.Status(e.Code).JSON(model.ApiResponse{
		Success: false,
		Error:   e,
		Meta:    response.GenerateMetadata(ctx),
	})
}

func ErrorBadRequest(err error) error {
	return NewApiError(
		err,
		fiber.StatusBadRequest,
		err.Error(),
		nil,
	)
}

func ErrorUnauthorized(err error, message string) error {
	return NewApiError(
		err,
		fiber.StatusUnauthorized,
		message,
		nil,
	)
}

func ErrorNotFound(err error) error {
	return NewApiError(
		err,
		fiber.StatusNotFound,
		"Resource not found",
		nil,
	)
}
