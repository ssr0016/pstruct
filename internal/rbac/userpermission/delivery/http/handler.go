package http

import (
	apiError "task-management-system/internal/api/errors"
	"task-management-system/internal/api/response"
	permission "task-management-system/internal/rbac/userpermission"

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

func (h *PermissionHandler) AddPermission(ctx *fiber.Ctx) error {
	var cmd permission.AddPermissionCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.AddPermission(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Created(ctx, fiber.Map{
		"permission": cmd,
	})
}

func (h *PermissionHandler) GetActionByUserID(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	result, err := h.s.GetActionByUserID(ctx.Context(), id)
	if err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"actions": result,
	})
}

func (h *PermissionHandler) GetListPermission(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	result, err := h.s.GetListPermission(ctx.Context(), id)
	if err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"permissions": result,
	})

}
