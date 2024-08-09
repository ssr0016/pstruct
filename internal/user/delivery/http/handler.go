package http

import (
	"task-management-system/internal/user"
	"task-management-system/internal/user/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Usecase *usecase.UserUsecase
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var u user.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := h.Usecase.CreateUser(&u); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(u)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	u, err := h.Usecase.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	return c.JSON(u)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var u user.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := h.Usecase.UpdateUser(&u); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(u)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := h.Usecase.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.Usecase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(users)
}
