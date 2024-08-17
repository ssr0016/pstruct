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

func (h *UserRolesHandler) AssignRoles(ctx *fiber.Ctx) error {
	var cmd userroles.CreateUserRolesCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	err := h.s.AssignRoles(ctx.Context(), &cmd)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"message": "User role assigned successfully",
	})
}

func (h *UserRolesHandler) RemoveUserRoles(ctx *fiber.Ctx) error {
	var cmd userroles.RemoveUserRolesCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	err := h.s.RemoveUserRoles(ctx.Context(), &cmd)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"message": "User role removed successfully",
	})
}

func (h *UserRolesHandler) GetUserRolesByID(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	roles, err := h.s.GetUserRolesByID(ctx.Context(), id)
	if err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"roles": roles,
	})

}

func (h *UserRolesHandler) UpdateUserRoles(ctx *fiber.Ctx) error {
	var cmd userroles.UpdateUserRolesCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.UpdateUserRoles(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"user roles": cmd,
	})
}

func (h *UserRolesHandler) SearchUserRoles(ctx *fiber.Ctx) error {
	var query userroles.SearchUserRolesQuery

	if err := ctx.QueryParser(&query); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	result, err := h.s.SearchUserRoles(ctx.Context(), &query)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"roles": result,
	})
}
