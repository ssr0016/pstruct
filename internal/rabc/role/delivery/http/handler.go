package http

import (
	apiError "task-management-system/internal/api/errors"
	"task-management-system/internal/api/response"
	rabc "task-management-system/internal/rabc/role"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	s rabc.Service
}

func NewRoleHandler(s rabc.Service) *RoleHandler {
	return &RoleHandler{
		s: s,
	}
}

func (h *RoleHandler) CreateRole(ctx *fiber.Ctx) error {
	var cmd rabc.CreateRoleCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.CreateRole(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Created(ctx, fiber.Map{
		"role": cmd,
	})
}

func (h *RoleHandler) GetRoleByID(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	result, err := h.s.GetRoleByID(ctx.Context(), id)
	if err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"role": result,
	})
}

func (h *RoleHandler) UpdateRole(ctx *fiber.Ctx) error {
	var cmd rabc.UpdateRoleCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.UpdateRole(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"role": cmd,
	})
}

func (h *RoleHandler) DeleteRole(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := h.s.DeleteRole(ctx.Context(), id); err != nil {
		if err.Error() == "role not found" {
			return apiError.ErrorNotFound(err)
		}
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"message": "Role deleted successfully",
	})
}

func (h *RoleHandler) SearchRole(ctx *fiber.Ctx) error {
	var query rabc.SearchRoleQuery

	if err := ctx.QueryParser(&query); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	result, err := h.s.SearchRole(ctx.Context(), &query)
	if err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"roles": result,
	})
}
