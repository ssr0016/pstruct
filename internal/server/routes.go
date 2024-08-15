package server

import (
	"context"
	"errors"
	"task-management-system/internal/api/response"
	"task-management-system/internal/db"
	departmentHttp "task-management-system/internal/department/delivery/http"
	"task-management-system/internal/middleware"
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
	dh *departmentHttp.DepartmentHandler,
) {

	api := s.app.Group("/api")
	api.Get("/", healthCheck(s.db))

	// User routes
	user := api.Group("/users")
	user.Post("/register", uh.CreateUser)
	user.Post("/login", uh.LoginUser)

	// Admin routes
	admin := api.Group("/admin")
	admin.Use(middleware.JWTProtected(s.jwtSecret))
	admin.Post("/users", uh.CreateUser)
	admin.Get("/users", uh.SearchUser)
	admin.Get("/users/:id", uh.GetUserByID)
	admin.Put("/users/:id", uh.UpdateUser)
	admin.Delete("/users/:id", uh.DeleteUser)

	// Department routes
	department := api.Group("/departments")
	department.Use(middleware.JWTProtected(s.jwtSecret))
	department.Post("/", dh.CreateDepartment)
	department.Get("/", dh.SearchDepartment)
	department.Get("/:id", dh.GetDepartmentByID)
	department.Put("/:id", dh.UpdateDepartment)
	department.Delete("/:id", dh.DeleteDepartment)

	// Task routes
	task := api.Group("/tasks")
	task.Post("/", th.CreateTask)
	task.Get("/", th.SearchTask)
	task.Get("/:id", th.GetTaskByID)
	task.Put("/:id", th.UpdateTask)
	task.Delete("/:id", th.DeleteTask)

}
