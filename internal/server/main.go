package server

import (
	"task-management-system/config"
	"task-management-system/internal/logger"
	taskHttp "task-management-system/internal/task/delivery/http"
	taskUsecase "task-management-system/internal/task/usecase"
	userHttp "task-management-system/internal/user/delivery/http"
	userUsecase "task-management-system/internal/user/usecase"

	"task-management-system/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	apiError "task-management-system/internal/api/errors"
)

type Server struct {
	app       *fiber.App
	port      string
	jwtSecret string
	db        db.DB
	cfg       *config.Config
	log       *logger.Logger
}

func NewServer(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: apiError.DefaultErrorHandler,
	})

	app.Use(cors.New())

	port := ":" + cfg.Port

	// Initialize SqlxDB with the provided database configuration
	sqlxDB := &db.SqlxDB{DB: cfg.DB}

	return &Server{
		app:       app,
		port:      port,
		jwtSecret: cfg.JwtSecret,
		db:        sqlxDB,
		cfg:       cfg,
		log:       cfg.Logger,
	}
}

func (s *Server) Start() error {
	ts := taskUsecase.NewTaskUsecase(s.db, s.cfg)
	th := taskHttp.NewTaskHandler(ts)

	uu := userUsecase.NewUserCase(s.db, s.cfg)
	uh := userHttp.NewUserHandler(uu)

	s.SetupRoutes(th, uh)
	return s.app.Listen(s.port)
}

func (s *Server) Stop() error {
	// s.db.Close()
	s.log.Sync()
	return s.app.Shutdown()
}
