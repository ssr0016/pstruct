package server

import (
	"context"
	"errors"
	"task-management-system/internal/api/response"
	"task-management-system/internal/db"
	departmentHttp "task-management-system/internal/department/delivery/http"
	"task-management-system/internal/middleware"
	permissionHttp "task-management-system/internal/rbac/permissions/delivery/http"
	permissionuserHttp "task-management-system/internal/rbac/permissionuser/delivery/http"
	"task-management-system/internal/rbac/permissionuser/repository/postgres"
	roleHttp "task-management-system/internal/rbac/role/delivery/http"
	userroleHttp "task-management-system/internal/rbac/userroles/delivery/http"
	taskHttp "task-management-system/internal/task/delivery/http"
	userHttp "task-management-system/internal/user/delivery/http"

	userroleRepo "task-management-system/internal/rbac/userroles/repository/postgres"

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
	rh *roleHttp.RoleHandler,
	ph *permissionHttp.PermissionHandler,
	urh *userroleHttp.UserRolesHandler,
	puuh *permissionuserHttp.PermissionUserHandler,

	permRepo *postgres.PermissionUserRepository,
	userRoleRepo *userroleRepo.UserRoleRepository,
) {

	api := s.app.Group("/api")
	api.Get("/", healthCheck(s.db))

	// Initialize PermissionsRepository

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

	// Role routes
	role := api.Group("/roles")
	role.Use(middleware.JWTProtected(s.jwtSecret))
	role.Post("/", rh.CreateRole)
	role.Get("/:id", rh.GetRoleByID)
	role.Get("/", rh.GetRoles)
	role.Delete("/:id", rh.DeleteRole)

	// UserRole routes
	userrole := api.Group("/userroles")
	admin.Use(middleware.JWTProtected(s.jwtSecret))
	userrole.Post("/", urh.AssignRoles)
	userrole.Get("/:id", urh.GetUserRolesByID)
	userrole.Delete("/:id", urh.RemoveUserRoles)

	// Permission routes
	permission := api.Group("/permissions")
	permission.Use(middleware.JWTProtected(s.jwtSecret))
	permission.Post("/", ph.CreatePermission)
	permission.Get("/", ph.GetUserPermissions)
	permission.Get("/:id", ph.GetPermissionByID)

	// PermissionUser routes
	permissionuser := api.Group("/permissionusers")
	permissionuser.Use(middleware.JWTProtected(s.jwtSecret))
	permissionuser.Post("/", puuh.CreaPermissionUser)
	permissionuser.Get("/", puuh.GetUsersPermissions)
	permissionuser.Get("/:id", puuh.GetUserPermissionByID)
	permissionuser.Delete("/:id", puuh.DeleteUserPermission)
	permissionuser.Get("user/:id", puuh.GetAllUserPermissions)

	// Department routes
	department := api.Group("/departments")
	admin.Use(middleware.JWTProtected(s.jwtSecret))
	department.Post("/", dh.CreateDepartment)
	department.Get("/", dh.SearchDepartment)
	department.Get("/:id", dh.GetDepartmentByID)
	department.Put("/:id", dh.UpdateDepartment)
	department.Delete("/:id", dh.DeleteDepartment)

	// Task routes
	task := api.Group("/tasks")
	task.Use(middleware.JWTProtected(s.jwtSecret))
	task.Use(middleware.RoleBasedAccessControl(userRoleRepo, "admin"))
	task.Post("/", middleware.PermissionMiddleware("create", permRepo), th.CreateTask)
	task.Get("/", middleware.PermissionMiddleware("read", permRepo), th.SearchTask)
	task.Get("/:id", middleware.PermissionMiddleware("read", permRepo), th.GetTaskByID)
	task.Put("/:id", middleware.PermissionMiddleware("update", permRepo), th.UpdateTask)
	task.Delete("/:id", middleware.PermissionMiddleware("delete", permRepo), th.DeleteTask)

	api.Get("/test-context", middleware.JWTProtected(s.jwtSecret), func(c *fiber.Ctx) error {
		userID := c.Locals("userID")
		userRoles := c.Locals("userRoles")
		return c.JSON(fiber.Map{
			"userID":    userID,
			"userRoles": userRoles,
		})
	})
}
