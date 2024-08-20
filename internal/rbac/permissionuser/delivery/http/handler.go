package http

import (
	apiError "task-management-system/internal/api/errors"
	"task-management-system/internal/api/response"
	"task-management-system/internal/rbac/permissionuser"

	"github.com/gofiber/fiber/v2"
)

type PermissionUserHandler struct {
	s permissionuser.Service
}

func NewPermissionUserHandler(s permissionuser.Service) *PermissionUserHandler {
	return &PermissionUserHandler{
		s: s,
	}
}

func (h *PermissionUserHandler) CreaPermissionUser(ctx *fiber.Ctx) error {
	var cmd permissionuser.CreateUserPermissionCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.CreateUserPermission(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"message": "user permission created successfully",
	})

}

func (h *PermissionUserHandler) GetUsersPermissions(ctx *fiber.Ctx) error {
	var query permissionuser.UserPermissionsQuery

	if err := ctx.QueryParser(&query); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	result, err := h.s.GetUsersPermissions(ctx.Context(), &query)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"permissions": result,
	})
}

func (h *PermissionUserHandler) DeleteUserPermission(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := h.s.DeleteUserPermission(ctx.Context(), id); err != nil {
		if err.Error() == "user permission not found" {
			return apiError.ErrorNotFound(err)
		}
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"message": "user permission deleted successfully",
	})
}

func (h *PermissionUserHandler) GetUserPermissionByID(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	result, err := h.s.GetUserPermissionByID(ctx.Context(), id)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, result)
}

func (h *PermissionUserHandler) GetAllUserPermissions(ctx *fiber.Ctx) error {
	// Extract 'id' from URL parameters as a string
	userID := ctx.Params("id")
	if userID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	// Call the use case method with the string userID
	result, err := h.s.GetAllUserPermissions(ctx.Context(), userID)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"permissions": result,
	})
}
