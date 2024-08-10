package server

import (
	"task-management-system/config"
	"task-management-system/internal/db"
	taskHttp "task-management-system/internal/task/delivery/http"
	taskUsecase "task-management-system/internal/task/usecase"
	userHttp "task-management-system/internal/user/delivery/http"
	userUsecase "task-management-system/internal/user/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"

	apiError "task-management-system/internal/api/error"
)

type Server struct {
	app       *fiber.App
	port      string
	jwtSecret string
	db        *sqlx.DB
	cfg       *config.Config
}

func NewServer(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: apiError.DefaultErrorHandler,
	})

	app.Use(cors.New())

	port := ":" + cfg.Port
	db := db.Connect(cfg.DatabaseUrl)

	return &Server{
		app:       app,
		port:      port,
		jwtSecret: cfg.JwtSecret,
		db:        db,
		cfg:       cfg,
	}
}

func (s *Server) Start() error {
	ts := taskUsecase.NewTaskUsecase(s.db, s.cfg)
	th := taskHttp.NewTaskHandler(ts)

	uu := userUsecase.NewUserCase(s.db)
	uh := userHttp.NewUserHandler(uu)

	s.SetupRoutes(th, uh)
	return s.app.Listen(s.port)
}

func (s *Server) Stop() error {
	s.db.Close()
	return s.app.Shutdown()
}
