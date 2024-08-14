package http

import (
	apiError "task-management-system/internal/api/errors"
	"task-management-system/internal/api/response"
	"task-management-system/internal/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	s user.Service
}

func NewUserHandler(s user.Service) *UserHandler {
	return &UserHandler{
		s: s,
	}
}

func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var cmd user.CreateUserRequest

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.CreateUser(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Created(ctx, fiber.Map{
		"user": cmd,
	})
}

func (h *UserHandler) LoginUser(ctx *fiber.Ctx) error {
	var cmd user.LoginUserRequest

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	result, err := h.s.GetUserByEmail(ctx.Context(), cmd.Email)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"user": result,
	})
}

func (h *UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	result, err := h.s.GetUserByID(ctx.Context(), id)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"user": result,
	})
}
