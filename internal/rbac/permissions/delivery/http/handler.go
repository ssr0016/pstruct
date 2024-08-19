package http

import (
	apiError "task-management-system/internal/api/errors"
	"task-management-system/internal/api/response"
	"task-management-system/internal/rbac/permissions"

	"github.com/gofiber/fiber/v2"
)

type PermissionHandler struct {
	s permissions.Service
}

func NewPermissionHandler(s permissions.Service) *PermissionHandler {
	return &PermissionHandler{
		s: s,
	}
}

func (h *PermissionHandler) CreatePermission(c *fiber.Ctx) error {
	var cmd permissions.CreatePermissionCommand

	if err := c.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.CreatePermissions(c.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Created(c, fiber.Map{
		"permission": cmd,
	})
}

func (h *PermissionHandler) GetUserPermissions(c *fiber.Ctx) error {
	result, err := h.s.GetPermissions(c.Context())
	if err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(c, fiber.Map{
		"permissions": result,
	})
}

func (h *PermissionHandler) GetPermissionByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	result, err := h.s.GetPermissionByID(c.Context(), id)
	if err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(c, result)
}
