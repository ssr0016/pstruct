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

	result, err := h.s.GetUserByEmail(ctx.Context(), &cmd)
	if err != nil {
		if err == user.ErrUserNotFound || err == user.ErrInvalidPassword {
			return apiError.ErrorUnauthorized(err, "Invalid email or password")
		}
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

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	var cmd user.UpdateUserRequest

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	err := h.s.UpdateUser(ctx.Context(), &cmd)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"updated user": cmd,
	})
}

func (h *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := h.s.DeleteUser(ctx.Context(), id); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"message": "User deleted successfully",
	})
}

func (h *UserHandler) SearchUser(ctx *fiber.Ctx) error {
	var query user.SearchUserQuery

	if err := ctx.QueryParser(&query); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	result, err := h.s.SearchUser(ctx.Context(), &query)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"users": result,
	})
}
