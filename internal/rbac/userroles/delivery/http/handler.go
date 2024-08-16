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

func (h *UserRolesHandler) Assign(ctx *fiber.Ctx) error {
	var ids userroles.UserRole

	if err := ctx.BodyParser(&ids); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := ids.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	err := h.s.Assign(ctx.Context(), ids.UserID, ids.RoleID)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"message": "user role assigned successfully",
		"user_id": ids.UserID,
		"role_id": ids.RoleID,
	})
}
