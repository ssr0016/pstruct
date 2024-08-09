package routes

import (
	"task-management-system/internal/task/delivery/http"
	"task-management-system/internal/task/repository/postgres"
	"task-management-system/internal/task/usecase"
	userHandler "task-management-system/internal/user/delivery/http"
	userRepo "task-management-system/internal/user/repository/postgres"
	userUsecase "task-management-system/internal/user/usecase"

	"task-management-system/internal/db"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Task repository, usecase, and handler initialization
	taskRepo := &postgres.TaskRepository{DB: db.DB}
	taskUsecase := &usecase.TaskUsecase{Repo: taskRepo}
	taskHandler := &http.TaskHandler{Usecase: taskUsecase}

	// User repository, usecase, and handler initialization
	userRepo := &userRepo.UserRepository{DB: db.DB}
	userUsecase := &userUsecase.UserUsecase{Repo: userRepo}
	userHandler := &userHandler.UserHandler{Usecase: userUsecase}

	// Task routes
	app.Post("/tasks", taskHandler.CreateTask)
	app.Get("/tasks/:id", taskHandler.GetTaskByID)
	app.Put("/tasks/:id", taskHandler.UpdateTask)
	app.Delete("/tasks/:id", taskHandler.DeleteTask)
	app.Get("/tasks", taskHandler.GetAllTasks)

	// User routes
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
	app.Get("/users", userHandler.GetAllUsers)
}
