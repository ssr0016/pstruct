package http

import (
	apiError "task-management-system/internal/api/errors"
	"task-management-system/internal/api/response"
	"task-management-system/internal/rbac/permission"

	"github.com/gofiber/fiber/v2"
)

type PermissionHandler struct {
	s permission.Service
}

func NewPermissionHandler(s permission.Service) *PermissionHandler {
	return &PermissionHandler{
		s: s,
	}
}

func (h *PermissionHandler) CreatePermission(ctx *fiber.Ctx) error {
	var cmd permission.CreatePermissionCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.CreatePermission(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Created(ctx, fiber.Map{
		"permission": cmd,
	})
}

func (h *PermissionHandler) GetPermissionByID(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	result, err := h.s.GetPermissionByID(ctx.Context(), id)
	if err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"permission": result,
	})
}

func (h *PermissionHandler) UpdatePermission(ctx *fiber.Ctx) error {
	var cmd permission.UpdatePermissionCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.UpdatePermission(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"permission": cmd,
	})
}

func (h *PermissionHandler) DeletePermission(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := h.s.DeletePermission(ctx.Context(), id); err != nil {
		if err.Error() == "permission not found" {
			return apiError.ErrorNotFound(err)
		}
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"message": "Permission deleted successfully",
	})
}

func (h *PermissionHandler) SearchPermission(ctx *fiber.Ctx) error {
	var query permission.SearchPermissionQuery

	if err := ctx.QueryParser(&query); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	result, err := h.s.SearchPermission(ctx.Context(), &query)
	if err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"permissions": result,
	})
}
