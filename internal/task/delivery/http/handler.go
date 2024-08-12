package http

import (
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

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var cmd task.CreateTaskCommand
	if err := c.BodyParser(&cmd); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := h.s.CreateTask(c.Context(), &cmd); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(cmd)
}

func (h *TaskHandler) GetTaskByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	t, err := h.s.GetTaskByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	return c.JSON(t)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	var cmd task.UpdateTaskCommand
	if err := c.BodyParser(&cmd); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := h.s.UpdateTask(c.Context(), &cmd); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(cmd)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := h.s.DeleteTask(c.Context(), id); err != nil {
		if err.Error() == "task not found" {
			return c.Status(fiber.StatusNotFound).SendString("Task not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Task deleted successfully")
}

func (h *TaskHandler) SearchTask(c *fiber.Ctx) error {
	var query task.SearchTaskQuery

	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	result, err := h.s.SearchTask(c.Context(), &query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(result)
}
