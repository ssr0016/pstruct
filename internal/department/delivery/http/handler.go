package http

import (
	apiError "task-management-system/internal/api/errors"
	"task-management-system/internal/api/response"
	"task-management-system/internal/department"

	"github.com/gofiber/fiber/v2"
)

type DepartmentHandler struct {
	s department.Service
}

func NewDepartmentHandler(s department.Service) *DepartmentHandler {
	return &DepartmentHandler{
		s: s,
	}
}

func (h *DepartmentHandler) CreateDepartment(ctx *fiber.Ctx) error {
	var cmd department.CreateDepartmentCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.CreateDepartment(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Created(ctx, fiber.Map{
		"department": cmd,
	})
}

func (h *DepartmentHandler) GetDepartmentByID(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	result, err := h.s.GetDepartmentByID(ctx.Context(), id)
	if err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"department": result,
	})
}

func (h *DepartmentHandler) UpdateDepartment(ctx *fiber.Ctx) error {
	var cmd department.UpdateDepartmentCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.UpdateDepartment(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"department": cmd,
	})
}

func (h *DepartmentHandler) DeleteDepartment(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := h.s.DeleteDepartment(ctx.Context(), id); err != nil {
		if err.Error() == "task not found" {
			return apiError.ErrorNotFound(err)
		}
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"message": "Department deleted successfully",
	})
}

func (h *DepartmentHandler) SearchDepartment(ctx *fiber.Ctx) error {
	var query department.SearchDepartmentQuery

	if err := ctx.QueryParser(&query); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	result, err := h.s.SearchDepartment(ctx.Context(), &query)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"departments": result,
	})
}
