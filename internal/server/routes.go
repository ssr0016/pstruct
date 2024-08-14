package server

import (
	"context"
	"errors"
	"task-management-system/internal/api/response"
	"task-management-system/internal/db"
	taskHttp "task-management-system/internal/task/delivery/http"
	userHttp "task-management-system/internal/user/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func healthCheck(db db.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var result int
		err := db.Get(context.Background(), &result, "SELECT 1")
		if err != nil {
			return errors.New("database unavailable")
		}
		return response.Ok(ctx, fiber.Map{
			"database": "available",
		})
	}
}

func (s *Server) SetupRoutes(
	th *taskHttp.TaskHandler,
	uh *userHttp.UserHandler,
) {

	api := s.app.Group("/api")
	api.Get("/", healthCheck(s.db))

	// Task routes
	task := api.Group("/tasks")
	task.Post("/", th.CreateTask)
	task.Get("/", th.SearchTask)
	task.Get("/:id", th.GetTaskByID)
	task.Put("/:id", th.UpdateTask)
	task.Delete("/:id", th.DeleteTask)

	// User routes
	// User routes
	user := api.Group("/users")
	user.Post("/", uh.CreateUser)
	user.Post("/login", uh.LoginUser)
	user.Get("/:id", uh.GetUserByID)

}
