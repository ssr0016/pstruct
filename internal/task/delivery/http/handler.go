package http

import (
	apiError "task-management-system/internal/api/errors"
	"task-management-system/internal/api/response"
	"task-management-system/internal/task"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	s task.Service
}

func NewTaskHandler(s task.Service) *TaskHandler {
	return &TaskHandler{
		s: s,
	}
}

func (h *TaskHandler) CreateTask(ctx *fiber.Ctx) error {
	var cmd task.CreateTaskCommand

	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.CreateTask(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Created(ctx, fiber.Map{
		"task": cmd,
	})
}

func (h *TaskHandler) GetTaskByID(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	result, err := h.s.GetTaskByID(ctx.Context(), id)
	if err != nil {
		return apiError.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"task": result,
	})
}

func (h *TaskHandler) UpdateTask(ctx *fiber.Ctx) error {
	var cmd task.UpdateTaskCommand
	if err := ctx.BodyParser(&cmd); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := cmd.Validate(); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	if err := h.s.UpdateTask(ctx.Context(), &cmd); err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"task": cmd,
	})
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := h.s.DeleteTask(c.Context(), id); err != nil {
		if err.Error() == "task not found" {
			return apiError.ErrorNotFound(err)
		}
		return apiError.ErrorInternalServerError(err)
	}

	return c.Status(fiber.StatusOK).SendString("Task deleted successfully")
}

func (h *TaskHandler) SearchTask(ctx *fiber.Ctx) error {
	var query task.SearchTaskQuery

	if err := ctx.QueryParser(&query); err != nil {
		return apiError.ErrorBadRequest(err)
	}

	result, err := h.s.SearchTask(ctx.Context(), &query)
	if err != nil {
		return apiError.ErrorInternalServerError(err)
	}

	return response.Ok(ctx, fiber.Map{
		"tasks": result,
	})
}
