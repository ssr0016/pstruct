package http

import (
	apiError "task-management-system/internal/api/errors"
	"task-management-system/internal/api/response"
	"task-management-system/internal/rbac/userroles"

	"github.com/gofiber/fiber/v2"
)

type UserRolesHandler struct {
	s userroles.Service
}

func NewUserRoleHandler(s userroles.Service) *UserRolesHandler {
	return &UserRolesHandler{
		s: s,
	}
}

func (h *UserRolesHandler) AssignRoleToUser(ctx *fiber.Ctx) error {
	var cmd userroles.CreateUserRoleCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.AssignRoleToUser(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"user": cmd,
	})
}

func (h *UserRolesHandler) RemoveRoleFromUser(ctx *fiber.Ctx) error {
	var cmd userroles.CreateRemoveUserRoleCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.RemoveRoleFromUser(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"user": cmd,
	})
}

func (h *UserRolesHandler) SearchUserRoles(ctx *fiber.Ctx) error {
	var query userroles.SearchUserRoleQuery

	if err := ctx.QueryParser(&query); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	result, err := h.s.SearchUserRole(ctx.Context(), &query)
	if err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"users": result,
	})
}
