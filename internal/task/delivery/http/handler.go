package http

import (
	"task-management-system/internal/task"
	"task-management-system/internal/task/usecase"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	s *usecase.TaskUseCase
}

func NewTaskHandler(s *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		s: s,
	}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var t task.Task
	if err := c.BodyParser(&t); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := h.s.CreateTask(&t); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(t)
}

func (h *TaskHandler) GetTaskByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	t, err := h.s.GetTaskByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	return c.JSON(t)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	var t task.Task
	if err := c.BodyParser(&t); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := h.s.UpdateTask(&t); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(t)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := h.s.DeleteTask(id); err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *TaskHandler) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := h.s.GetAllTasks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(tasks)
}
